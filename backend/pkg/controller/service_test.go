package controller_test

import (
	"log"
	"os"
	"testing"

	s "github.com/adventar/adventar/backend/pkg/controller"
	"github.com/adventar/adventar/backend/pkg/infra"
	db_client "github.com/adventar/adventar/backend/pkg/infra/db"
	"github.com/adventar/adventar/backend/pkg/usecase"
	"github.com/adventar/adventar/backend/pkg/util"
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
	dsn := "root@tcp(127.0.0.1:13306)/adventar_test?parseTime=true"
	db, err = sqlx.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	dbClient, err := db_client.New(dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer dbClient.Close()
	clients := infra.New(infra.WithDB(dbClient))
	f := &testMetaFetcher{}
	usecase := usecase.New(clients, f)

	v := &testVerifier{}
	service = s.NewService(db, v, usecase, clients)
	code := m.Run()
	os.Exit(code)
}

func cleanupDatabase() {
	var names []string
	err := db.Select(&names, "show tables")

	if err != nil {
		log.Fatal(err)
	}
	if _, err := db.Exec("set foreign_key_checks = 0"); err != nil {
		log.Fatal(err)
	}
	for _, name := range names {
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
	res, err := db.Exec(
		"insert into users (name, auth_uid, auth_provider, icon_url) values (?, ?, ?, ?)",
		u.name, u.authUID, u.authProvider, u.iconURL,
	)
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
	res, err := db.Exec(
		"insert into calendars(user_id, title, description, year) values(?, ?, ?, ?)",
		c.userID, c.title, c.description, c.year,
	)
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
	res, err := db.Exec(
		"insert into entries(user_id, calendar_id, day, url, comment, title, image_url) values(?, ?, ?, ?, ?, ?, ?)",
		e.userID, e.calendarID, e.day, e.url, e.comment, e.title, e.imageURL,
	)
	if err != nil {
		t.Fatal(err)
	}

	e.id, err = res.LastInsertId()
	if err != nil {
		t.Fatal(err)
	}
}
