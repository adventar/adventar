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
)

var (
	db      *sql.DB
	service *main.Service
)

func TestMain(m *testing.M) {
	var err error
	db, err = sql.Open("mysql", "root@tcp(127.0.0.1:3306)/adventar_test")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	service = main.NewService(db)
	code := m.Run()
	os.Exit(code)
}

func TestGetCalendar(t *testing.T) {
	cleanupDatabase()

	in := new(pb.GetCalendarRequest)
	ctx := context.Background()

	id, err := createCalendar()
	if err != nil {
		t.Fatal(err)
	}
	in.Id = id

	calendar, err := service.GetCalendar(ctx, in)
	if err != nil {
		t.Fatal(err)
	}

	if calendar.Id != in.Id {
		t.Errorf("actual: %d, expected: %d", calendar.Id, in.Id)
	}
}

func TestCreateCalendar(t *testing.T) {
	cleanupDatabase()

	userID, err := createUser()
	if err != nil {
		t.Fatal(err)
	}
	in := &pb.CreateCalendarRequest{UserId: userID, Title: "foo", Description: "bar", Year: 2019}
	ctx := context.Background()

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

func createUser() (int64, error) {
	stmt, err := db.Prepare("insert into users (name, auth_uid, auth_provider, icon_url) values (?, ?, ?, ?)")
	defer stmt.Close()

	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec("test user", "xxx", "google", "")
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

func createCalendar() (int64, error) {
	userID, err := createUser()
	if err != nil {
		return 0, err
	}

	stmt, err := db.Prepare("insert into calendars(user_id, title, description, year) values(?, ?, ?, ?)")
	defer stmt.Close()

	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(userID, "test title", "test description", 2019)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}
