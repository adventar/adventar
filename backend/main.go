package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	adventarv1 "github.com/adventar/adventar/backend/pkg/gen/adventar/v1"
	"github.com/adventar/adventar/backend/pkg/gen/adventar/v1/adventarv1connect"
	"github.com/adventar/adventar/backend/pkg/model"
	"github.com/bufbuild/connect-go"
	"github.com/jmoiron/sqlx"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type AdventarServer struct {
	// TODO: remove
	adventarv1connect.UnimplementedAdventarHandler
	db *sqlx.DB
}

func (s *AdventarServer) GetCalendar(
	ctx context.Context,
	req *connect.Request[adventarv1.GetCalendarRequest],
) (*connect.Response[adventarv1.GetCalendarResponse], error) {
	id := req.Msg.CalendarId
	calendar, err := getCalendar(s.db, id)
	if err != nil {
		log.Printf("Error: %v", err)
	}
	res := connect.NewResponse(&adventarv1.GetCalendarResponse{
		Calendar: &adventarv1.Calendar{
			Id:    calendar.ID,
			Title: calendar.Title,
			Year:  calendar.Year,
		},
	})
	return res, nil
}

func main() {
	db := setupDatabase()
	defer db.Close()
	mux := http.NewServeMux()
	mux.Handle(adventarv1connect.NewAdventarHandler(&AdventarServer{}))
	err := http.ListenAndServe(
		"localhost:8080",
		h2c.NewHandler(mux, &http2.Server{}),
	)
	log.Fatalf("listen failed: %v", err)
}

func setupDatabase() *sqlx.DB {
	var err error
	source := os.Getenv("DATABASE_SOURCE")
	if source == "" {
		source = "root@tcp(127.0.0.1:13306)/adventar_dev"
	}
	source += "?parseTime=true&charset=utf8mb4"
	db, err := sqlx.Open("mysql", source)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func getCalendar(db *sqlx.DB, id int64) (*model.Calendar, error) {
	query := "select id, title, year from calendars where id = ?"
	var calendar model.Calendar
	err := db.Get(&calendar, query, id)

	if err != nil {
		return nil, fmt.Errorf("error: %v", err)
	}

	return &calendar, nil
}
