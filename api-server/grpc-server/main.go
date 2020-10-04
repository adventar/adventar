package main

import (
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"github.com/adventar/adventar/api-server/grpc-server/service"
	"github.com/adventar/adventar/api-server/grpc-server/util"
)

func main() {
	var err error
	source := os.Getenv("DATABASE_SOURCE")
	if source == "" {
		source = "root@tcp(127.0.0.1:13306)/adventar_dev?parseTime=true"
	}
	db, err := sqlx.Open("mysql", source)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	v := &util.FirebaseVerifier{}
	f := &util.SiteMetaFetcher{}
	s := service.NewService(db, v, f)
	s.Serve(":8080")
}
