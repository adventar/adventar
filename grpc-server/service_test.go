package main_test

import (
	"context"
	"database/sql"
	"log"
	"os"
	"testing"

	main "github.com/adventar/adventar/grpc-server"
	pb "github.com/adventar/adventar/grpc-server/adventar/v1"
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc/metadata"
)

var (
	db      *sql.DB
	service *main.Service
)

type testVerifier struct{}

func (v *testVerifier) VerifyIDToken(s string) *main.AuthResult {
	return &main.AuthResult{
		Name:         "foo",
		IconURL:      "http://example.com/icon",
		AuthProvider: "google",
		AuthUID:      "xxx",
	}
}

func TestMain(m *testing.M) {
	var err error
	db, err = sql.Open("mysql", "root@tcp(127.0.0.1:3306)/adventar_test")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	v := &testVerifier{}
	service = main.NewService(db, v)
	code := m.Run()
	os.Exit(code)
}

func TestListCalendars(t *testing.T) {
	cleanupDatabase()

	in := &pb.ListCalendarsRequest{Year: 2019}
	ctx := context.Background()

	u := &user{name: "foo", authUID: "xxx", authProvider: "google"}
	createUser(t, u)

	c := &calendar{title: "a", description: "b", userID: u.id, year: 2019}
	createCalendar(t, c)

	e := &entry{userID: u.id, calendarID: c.id, date: "2019-12-01"}
	createEntry(t, e)

	res, err := service.ListCalendars(ctx, in)
	if err != nil {
		t.Fatal(err)
	}

	if len(res.Calendars) != 1 {
		t.Errorf("actual: %d, expected: %d", len(res.Calendars), 1)
	}
	if res.Calendars[0].GetOwner().GetId() != u.id {
		t.Errorf("actual: %d, expected: %d", res.Calendars[0].GetEntryCount(), 1)
	}
	if res.Calendars[0].GetEntryCount() != 1 {
		t.Errorf("actual: %d, expected: %d", res.Calendars[0].GetEntryCount(), 1)
	}
}

func TestListCalendarsWithQuery(t *testing.T) {
	cleanupDatabase()

	in := &pb.ListCalendarsRequest{Query: "test", Year: 2019}
	ctx := context.Background()

	u := &user{name: "foo", authUID: "xxx", authProvider: "google"}
	createUser(t, u)

	c1 := &calendar{title: "Test title", description: "", userID: u.id, year: 2019}
	createCalendar(t, c1)

	c2 := &calendar{title: "foo", description: "Calendar test", userID: u.id, year: 2019}
	createCalendar(t, c2)

	c3 := &calendar{title: "foo", description: "bar", userID: u.id, year: 2019}
	createCalendar(t, c3)

	res, err := service.ListCalendars(ctx, in)
	if err != nil {
		t.Fatal(err)
	}

	if len(res.Calendars) != 2 {
		t.Errorf("actual: %d, expected: %d", len(res.Calendars), 2)
	}
}

func TestListCalendarsWithUserId(t *testing.T) {
	cleanupDatabase()

	users := []*user{}
	for _, n := range []string{"u1", "u2", "u3"} {
		u := &user{name: n, authUID: n}
		createUser(t, u)
		c := &calendar{userID: u.id, year: 2019}
		createCalendar(t, c)
		users = append(users, u)
	}

	in := &pb.ListCalendarsRequest{UserId: users[0].id, Year: 2019}
	ctx := context.Background()

	res, err := service.ListCalendars(ctx, in)
	if err != nil {
		t.Fatal(err)
	}

	if len(res.Calendars) != 1 {
		t.Errorf("actual: %d, expected: %d", len(res.Calendars), 1)
	}
}

func TestGetCalendar(t *testing.T) {
	cleanupDatabase()

	in := new(pb.GetCalendarRequest)
	ctx := context.Background()

	u := &user{name: "foo", authUID: "xxx", authProvider: "google"}
	createUser(t, u)

	c := &calendar{title: "a", description: "b", userID: u.id, year: 2019}
	createCalendar(t, c)

	e := &entry{userID: u.id, calendarID: c.id, date: "2019-12-01"}
	createEntry(t, e)

	in.CalendarId = c.id

	res, err := service.GetCalendar(ctx, in)
	if err != nil {
		t.Fatal(err)
	}

	if res.Calendar.Id != in.CalendarId {
		t.Errorf("actual: %d, expected: %d", res.Calendar.Id, in.CalendarId)
	}

	if res.Calendar.EntryCount != 1 {
		t.Errorf("actual: %d, expected: %d", res.Calendar.EntryCount, 1)
	}

	if res.Entries[0].Id != e.id {
		t.Errorf("actual: %d, expected: %d", res.Entries[0].Id, e.id)
	}
}

func TestCreateCalendar(t *testing.T) {
	cleanupDatabase()

	u := &user{name: "foo", authUID: "xxx", authProvider: "google"}
	createUser(t, u)

	in := &pb.CreateCalendarRequest{Title: "foo", Description: "bar"}
	md := make(map[string][]string)
	md["authorization"] = append(md["authorization"], "x")
	ctx := metadata.NewIncomingContext(context.Background(), md)

	calendar, err := service.CreateCalendar(ctx, in)
	if err != nil {
		t.Fatal(err)
	}

	var count int
	db.QueryRow("select count(*) from calendars").Scan(&count)
	if count != 1 {
		t.Errorf("actual: %d, expected: 1", count)
	}

	var title string
	db.QueryRow("select title from calendars where id = ?", calendar.Id).Scan(&title)
	if title != "foo" {
		t.Errorf("actual: %s, expected: foo", title)
	}
}

func TestSignInIfUserExists(t *testing.T) {
	cleanupDatabase()

	u := &user{name: "foo", authUID: "xxx", authProvider: "google"}
	createUser(t, u)

	in := &pb.SignInRequest{Jwt: ""}
	ctx := context.Background()
	_, err := service.SignIn(ctx, in)
	if err != nil {
		t.Fatal(err)
	}
	var iconURL string
	err = db.QueryRow("select icon_url from users").Scan(&iconURL)
	if err != nil {
		t.Fatal(err)
	}
	expected := "http://example.com/icon"
	if iconURL != expected {
		t.Errorf("actual: %s, expected: %s", iconURL, expected)
	}
}

func TestSignInIfUserDoesNotExist(t *testing.T) {
	cleanupDatabase()
	in := &pb.SignInRequest{Jwt: ""}
	ctx := context.Background()
	_, err := service.SignIn(ctx, in)
	if err != nil {
		t.Fatal(err)
	}
	var count int
	err = db.QueryRow("select count(*) users").Scan(&count)
	if err != nil {
		t.Fatal(err)
	}
	if count != 1 {
		t.Errorf("actual: %d, expected: 1", count)
	}
}

func TestUpdateUser(t *testing.T) {
	cleanupDatabase()

	u := &user{name: "foo", authUID: "xxx", authProvider: "google"}
	createUser(t, u)

	in := &pb.UpdateUserRequest{Name: "changed"}
	md := make(map[string][]string)
	md["authorization"] = append(md["authorization"], "x")
	ctx := metadata.NewIncomingContext(context.Background(), md)

	_, err := service.UpdateUser(ctx, in)
	if err != nil {
		t.Fatal(err)
	}

	var name string
	err = db.QueryRow("select name from users where id = ?", u.id).Scan(&name)
	if err != nil {
		t.Fatal(err)
	}

	if name != "changed" {
		t.Errorf("actual: %s, expected: %s", name, "changed")
	}
}

func cleanupDatabase() {
	rows, err := db.Query("show tables")

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	if _, err := db.Exec("set foreign_key_checks = 0"); err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			log.Fatal(err)
		}
		if _, err := db.Exec("truncate table " + name); err != nil {
			log.Fatal(err)
		}
	}
	if _, err := db.Exec("set foreign_key_checks = 1"); err != nil {
		log.Fatal(err)
	}
}

type user struct {
	id           int64
	name         string
	authUID      string
	authProvider string
	iconURL      string
}

func createUser(t *testing.T, u *user) {
	stmt, err := db.Prepare("insert into users (name, auth_uid, auth_provider, icon_url) values (?, ?, ?, ?)")
	defer stmt.Close()

	if err != nil {
		t.Fatal(err)
	}

	res, err := stmt.Exec(u.name, u.authUID, u.authProvider, u.iconURL)
	if err != nil {
		t.Fatal(err)
	}

	u.id, err = res.LastInsertId()
	if err != nil {
		t.Fatal(err)
	}
}

type calendar struct {
	id          int64
	userID      int64
	title       string
	description string
	year        int
}

func createCalendar(t *testing.T, c *calendar) {
	stmt, err := db.Prepare("insert into calendars(user_id, title, description, year) values(?, ?, ?, ?)")
	defer stmt.Close()

	if err != nil {
		t.Fatal(err)
	}

	res, err := stmt.Exec(c.userID, c.title, c.description, c.year)
	if err != nil {
		t.Fatal(err)
	}

	c.id, err = res.LastInsertId()
	if err != nil {
		t.Fatal(err)
	}
}

type entry struct {
	id         int64
	calendarID int64
	userID     int64
	date       string
	url        string
	comment    string
	title      string
	imageURL   string
}

func createEntry(t *testing.T, e *entry) {
	stmt, err := db.Prepare("insert into entries(user_id, calendar_id, date, url, comment, title, image_url) values(?, ?, ?, ?, ?, ?, ?)")
	defer stmt.Close()

	if err != nil {
		t.Fatal(err)
	}

	res, err := stmt.Exec(e.userID, e.calendarID, e.date, e.url, e.comment, e.title, e.imageURL)
	if err != nil {
		t.Fatal(err)
	}

	e.id, err = res.LastInsertId()
	if err != nil {
		t.Fatal(err)
	}
}
