package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var err error
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/adventar_dev")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	v := &firebaseVerifier{}
	f := &siteMetaFetcher{}
	s := NewService(db, v, f)
	s.serve(":8080")
}
