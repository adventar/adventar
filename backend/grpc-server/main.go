package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/adventar/adventar/backend/grpc-server/service"
	"github.com/adventar/adventar/backend/grpc-server/util"
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

	v := &util.FirebaseVerifier{}
	f := &util.SiteMetaFetcher{}
	s := service.NewService(db, v, f)
	s.Serve(":8080")
}
