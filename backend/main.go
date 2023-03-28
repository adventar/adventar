package main

import (
	"os"

	"github.com/adventar/adventar/backend/pkg/infra"
	"github.com/adventar/adventar/backend/pkg/infra/db"
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
	db_, err := sqlx.Open("mysql", source)
	if err != nil {
		util.Logger.Fatal().Err(err).Msg("")
	}
	defer db_.Close()

	dbClient, err := db.New(source)
	if err != nil {
		util.Logger.Fatal().Err(err).Msg("")
	}
	defer dbClient.Close()
	clients := infra.New(infra.WithDB(dbClient))

	v := &util.FirebaseVerifier{}
	f := &util.SiteMetaFetcher{}
	s := service.NewService(db_, v, f, clients)
	s.Serve(":8080")
}
