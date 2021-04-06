package main

import (
	"log"
	"os"
	"time"

	util "github.com/adventar/adventar/api-server/grpc-server/util"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Entry struct {
	ID  int64  `db:"id"`
	URL string `db:"url"`
}

func main() {
	err := run()

	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	var err error
	source := os.Getenv("DATABASE_SOURCE")
	source += "?parseTime=true&charset=utf8mb4"
	db, err := sqlx.Open("mysql", source)
	if err != nil {
		return err
	}
	defer db.Close()

	args := os.Args[1:]
	if len(args) == 0 {
		return UpdateTodaysEntries(db)
	} else {
		return UpdateEntryByIds(db, args)
	}
}

func updateEntries(db *sqlx.DB, entries []Entry) error {
	fetcher := &util.SiteMetaFetcher{}

	for i, entry := range entries {
		log.Printf("[Process] id: %d, url: %s", entry.ID, entry.URL)
		meta, err := fetcher.Fetch(entry.URL)
		if err != nil {
			log.Printf("Fetch error: %s", err)
			continue
		}

		_, err = db.Exec("update entries set title = ?, image_url = ? where id = ?", meta.Title, meta.ImageURL, entry.ID)
		if err != nil {
			log.Printf("Update error: %s", err)
			continue
		}

		log.Printf("[Updated %d] id: %d, title: %s, image_url: %s", i, entry.ID, meta.Title, meta.ImageURL)
	}

	log.Printf("Finish, processed count: %d", len(entries))

	return nil
}

func UpdateTodaysEntries(db *sqlx.DB) error {
	now, err := util.CurrentDate()
	if err != nil {
		return err
	}

	log.Printf("UpdateTodaysEntries")
	log.Printf("now: %v", now)
	log.Printf("time: %v", time.Now())

	var entries []Entry
	err = db.Select(&entries, `
		select
			e.id,
			e.url
		from
			entries as e
			inner join calendars as c on e.calendar_id = c.id
		where
			c.year = ? and
			e.day = ? and
			e.url != "" and
			e.title = "" and
			e.image_url = ""
		`, now.Year, now.Day)
	if err != nil {
		return err
	}

	return updateEntries(db, entries)
}

func UpdateEntryByIds(db *sqlx.DB, ids []string) error {
	log.Printf("UpdateEntryByIds")
	log.Printf("ids: %v", ids)

	var entries []Entry
	sql, args, err := sqlx.In("select id, url from entries where id in (?) and url != ''", ids)
	if err != nil {
		return err
	}
	err = db.Select(&entries, sql, args...)
	if err != nil {
		return err
	}

	return updateEntries(db, entries)
}
