package main

import (
	"database/sql"
	"log"
	"os"
	"time"

	util "github.com/adventar/adventar/api-server/grpc-server/util"
	_ "github.com/go-sql-driver/mysql"
)

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
	db, err := sql.Open("mysql", source)
	if err != nil {
		return err
	}
	defer db.Close()

	fetcher := &util.SiteMetaFetcher{}
	now, err := util.CurrentDate()
	if err != nil {
		return err
	}
	log.Printf("now: %v", now)
	log.Printf("time: %v", time.Now())

	rows, err := db.Query(`
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
	defer rows.Close()

	count := 0
	for rows.Next() {
		var id int
		var url string
		err := rows.Scan(&id, &url)
		if err != nil {
			return err
		}

		log.Printf("[Process] id: %d, url: %s", id, url)
		meta, err := fetcher.Fetch(url)
		if err != nil {
			log.Printf("Fetch error: %s", err)
			continue
		}

		stmt, err := db.Prepare("update entries set title = ?, image_url = ? where id = ?")
		if err != nil {
			log.Printf("Update error: %s", err)
			continue
		}
		defer stmt.Close()
		_, err = stmt.Exec(meta.Title, meta.ImageURL, id)
		if err != nil {
			log.Printf("Query execution error: %s", err)
			continue
		}

		count++
		log.Printf("[Updated] id: %d, title: %s, image_url: %s", id, meta.Title, meta.ImageURL)
	}

	log.Printf("Finish, processed count: %d", count)

	return nil
}
