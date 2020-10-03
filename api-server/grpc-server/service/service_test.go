package service_test

import (
	"log"
	"os"
	"testing"

	s "github.com/adventar/adventar/api-server/grpc-server/service"
	"github.com/adventar/adventar/api-server/grpc-server/util"
	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

type testVerifier struct{}

func (v *testVerifier) VerifyIDToken(s string) (*util.AuthResult, error) {
	return &util.AuthResult{
		Name:         "foo",
		IconURL:      "http://example.com/icon",
		AuthProvider: "google",
		AuthUID:      s,
	}, nil
}

type testMetaFetcher struct{}

func (tmf *testMetaFetcher) Fetch(url string) (*util.SiteMeta, error) {
	return &util.SiteMeta{Title: "site title", ImageURL: "http://example.com/image"}, nil
}

var (
	db      *sqlx.DB
	service *s.Service
)

func TestMain(m *testing.M) {
	var err error
	db, err = sqlx.Open("mysql", "root@tcp(127.0.0.1:3306)/adventar_test")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	v := &testVerifier{}
	f := &testMetaFetcher{}
	service = s.NewService(db, v, f)
	code := m.Run()
	os.Exit(code)
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
	day        int32
	url        string
	comment    string
	title      string
	imageURL   string
}

func createEntry(t *testing.T, e *entry) {
	stmt, err := db.Prepare("insert into entries(user_id, calendar_id, day, url, comment, title, image_url) values(?, ?, ?, ?, ?, ?, ?)")
	defer stmt.Close()

	if err != nil {
		t.Fatal(err)
	}

	res, err := stmt.Exec(e.userID, e.calendarID, e.day, e.url, e.comment, e.title, e.imageURL)
	if err != nil {
		t.Fatal(err)
	}

	e.id, err = res.LastInsertId()
	if err != nil {
		t.Fatal(err)
	}
}
