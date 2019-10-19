package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var err error
	source := os.Getenv("DATABASE_SOURCE")
	if source == "" {
		source = "root@tcp(127.0.0.1:13306)/adventar_dev"
	}
	db, err := sql.Open("mysql", source)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	v := &firebaseVerifier{}
	f := &siteMetaFetcher{}
	s := NewService(db, v, f)
	s.serve(":8080")
}
