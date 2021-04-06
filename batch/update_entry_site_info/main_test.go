package main_test

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"

	m "github.com/adventar/adventar/batch/update_entry_site_info"
)

type Entry struct {
	ID         int64     `db:"id"`
	UserID     int64     `db:"user_id"`
	CalendarID int64     `db:"calendar_id"`
	Day        int32     `db:"day"`
	Comment    string    `db:"comment"`
	URL        string    `db:"url"`
	Title      string    `db:"title"`
	ImageURL   string    `db:"image_url"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}

type input struct {
	entry    *Entry
	title    string
	imageUrl string
}

func assertEntries(t *testing.T, db *sqlx.DB, in []input) {
	for _, data := range in {
		var entry Entry
		err := db.Get(&entry, "select * from entries where id = ?", data.entry.ID)
		if err != nil {
			t.Fatal(err)
		}
		if entry.Title != data.title {
			t.Errorf("actual: %s, expected: %s", entry.Title, data.title)
		}
		if entry.ImageURL != data.imageUrl {
			t.Errorf("actual: %s, expected: %s", entry.ImageURL, data.imageUrl)
		}
	}
}

func TestUpdateByIds(t *testing.T) {
	db, ts := initTesting()
	defer db.Close()
	defer ts.Close()

	cid1 := createCalendar(db, 2020)
	entry1 := createEntry(db, &Entry{CalendarID: cid1, Day: 1, Title: "", ImageURL: "", URL: ts.URL})
	entry2 := createEntry(db, &Entry{CalendarID: cid1, Day: 2, Title: "", ImageURL: "", URL: ""})
	entry3 := createEntry(db, &Entry{CalendarID: cid1, Day: 3, Title: "a", ImageURL: "b", URL: ts.URL})

	err := m.UpdateEntryByIds(db, []string{
		strconv.FormatInt(entry1.ID, 10),
		strconv.FormatInt(entry2.ID, 10),
		strconv.FormatInt(entry3.ID, 10),
	})
	if err != nil {
		t.Fatal(err)
	}

	assertEntries(t, db, []input{
		{entry1, "foo", "bar"},
		{entry2, "", ""},
		{entry3, "foo", "bar"},
	})
}

func TestUpdateTodaysEntries(t *testing.T) {
	initTesting()
	db, ts := initTesting()
	defer db.Close()
	defer ts.Close()

	cid1 := createCalendar(db, 2020)
	entry1 := createEntry(db, &Entry{CalendarID: cid1, Day: 1, Title: "", ImageURL: "", URL: ts.URL})
	entry2 := createEntry(db, &Entry{CalendarID: cid1, Day: 2, Title: "", ImageURL: "", URL: ""})
	entry3 := createEntry(db, &Entry{CalendarID: cid1, Day: 3, Title: "a", ImageURL: "b", URL: ts.URL})

	cid2 := createCalendar(db, 2020)
	entry4 := createEntry(db, &Entry{CalendarID: cid2, Day: 1, Title: "", ImageURL: "", URL: ts.URL})
	entry5 := createEntry(db, &Entry{CalendarID: cid2, Day: 2, Title: "", ImageURL: "", URL: ts.URL})

	os.Setenv("CURRENT_DATE", "2020-12-01 10:00:00")
	err := m.UpdateTodaysEntries(db)
	if err != nil {
		t.Fatal(err)
	}
	assertEntries(t, db, []input{
		{entry1, "foo", "bar"},
		{entry2, "", ""},
		{entry3, "a", "b"},
		{entry4, "foo", "bar"},
		{entry5, "", ""},
	})

	os.Setenv("CURRENT_DATE", "2020-12-02 10:00:00")
	err = m.UpdateTodaysEntries(db)
	if err != nil {
		t.Fatal(err)
	}
	assertEntries(t, db, []input{
		{entry2, "", ""},
		{entry5, "foo", "bar"},
	})

	os.Setenv("CURRENT_DATE", "2020-12-03 10:00:00")
	err = m.UpdateTodaysEntries(db)
	if err != nil {
		t.Fatal(err)
	}
	assertEntries(t, db, []input{
		{entry3, "a", "b"},
	})
}

func initTesting() (*sqlx.DB, *httptest.Server) {
	log.SetFlags(log.Lshortfile)
	db, err := sqlx.Open("mysql", "root@tcp(127.0.0.1:3306)/adventar_test?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	var names []string
	err = db.Select(&names, "show tables")

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
	if _, err := db.Exec("insert into users (id, name, auth_uid, auth_provider, icon_url) values (1, 'x', 'y', 'z', '')"); err != nil {
		log.Fatal(err)
	}

	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `<head><title>foo</title><meta property="og:image" content="bar"><body>`)
	})
	ts := httptest.NewServer(h)

	return db, ts
}

func createCalendar(db *sqlx.DB, year int) int64 {
	res, err := db.Exec("insert into calendars(user_id, title, description, year) values(1, 't', 'd', ?)", year)
	if err != nil {
		log.Fatal(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	return id
}

func createEntry(db *sqlx.DB, e *Entry) *Entry {
	res, err := db.Exec(
		"insert into entries(user_id, calendar_id, day, url, title, image_url, comment) values(1, ?, ?, ?, ?, ?, '')",
		e.CalendarID, e.Day, e.URL, e.Title, e.ImageURL,
	)
	if err != nil {
		log.Fatal(err)
	}

	e.ID, err = res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	return e
}
