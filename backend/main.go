package main

import (
	"os"

	"github.com/adventar/adventar/backend/pkg/controller"
	"github.com/adventar/adventar/backend/pkg/infra"
	"github.com/adventar/adventar/backend/pkg/infra/db"
	"github.com/adventar/adventar/backend/pkg/usecase"
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
		panic(err)
	}
	defer db_.Close()

	dbClient, err := db.New(source)
	if err != nil {
		panic(err)
	}
	defer dbClient.Close()
	v := &util.FirebaseVerifier{}
	metaFetcher := &util.SiteMetaFetcher{}
	clients := infra.New(infra.WithDB(dbClient))
	usecase := usecase.New(clients, metaFetcher)

	s := controller.NewService(db_, v, usecase, clients)
	s.Serve(":8080")
}
