package main

import (
	"os"

	"github.com/adventar/adventar/backend/pkg/service"
	"github.com/adventar/adventar/backend/pkg/util"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {
	var err error
	source := os.Getenv("DATABASE_SOURCE")
	if source == "" {
		source = "root@tcp(127.0.0.1:13306)/adventar_dev"
	}
	source += "?parseTime=true&charset=utf8mb4"
	db, err := sqlx.Open("mysql", source)
	if err != nil {
		util.Logger.Fatal().Err(err).Msg("")
	}
	defer db.Close()

	v := &util.FirebaseVerifier{}
	f := &util.SiteMetaFetcher{}
	s := service.NewService(db, v, f)
	s.Serve(":8080")
}
