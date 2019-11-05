package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	util "github.com/adventar/adventar/backend/grpc-server/util"
	"github.com/aws/aws-lambda-go/lambda"
	_ "github.com/go-sql-driver/mysql"
)

func handler(ctx context.Context) error {
	var err error
	source := os.Getenv("DATABASE_SOURCE")
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
	fmt.Printf("now: %v", now)

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
	for rows.Next() {
		var id int
		var url string
		err := rows.Scan(&id, &url)
		if err != nil {
			return err
		}

		fmt.Printf("[Process] id: %d, url: %s", id, url)
		meta, err := fetcher.Fetch(url)
		if err != nil {
			fmt.Printf("Fetch error: %s", err)
		}

		stmt, err := db.Prepare("update entries set title = ?, image_url = ? where id = ?")
		if err != nil {
			return err
		}
		defer stmt.Close()
		_, err = stmt.Exec(meta.Title, meta.ImageURL, id)
		if err != nil {
			fmt.Printf("Query execution: %s", err)
		}

		fmt.Printf("[Updated] id: %d, title: %s, image_url: %s", id, meta.Title, meta.ImageURL)
	}

	return nil
}

func main() {
	lambda.Start(handler)
}
