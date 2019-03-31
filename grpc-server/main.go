package main

import (
	"context"
	"database/sql"
	"log"
	"net"

	pb "github.com/adventar/adventar/grpc-server/adventar/v1"
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

var db *sql.DB

type server struct{}

type Calendar struct {
	ID          int64
	UserID      int64
	Title       string
	Description string
	Year        int32
}

func (s *server) GetCalendar(ctx context.Context, in *pb.GetCalendarRequest) (*pb.Calendar, error) {
	var calendar Calendar
	log.Printf("Request Id: %d", in.GetId())
	err := db.QueryRow("select id, user_id, title, description, year from calendars where id = ?", in.GetId()).Scan(&calendar.ID, &calendar.UserID, &calendar.Title, &calendar.Description, &calendar.Year)
	if err != nil {
		return nil, err
	}
	return &pb.Calendar{Id: calendar.ID, UserId: calendar.UserID, Title: calendar.Title, Description: calendar.Description, Year: calendar.Year}, nil
}

func (s *server) CreateCalendar(ctx context.Context, in *pb.CreateCalendarRequest) (*pb.Calendar, error) {
	stmt, err := db.Prepare("insert into calendars(user_id, title, description, year) values(?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}

	res, err := stmt.Exec(in.GetUserId(), in.GetTitle(), in.GetDescription(), in.GetYear())
	if err != nil {
		return nil, err
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	var calendar Calendar
	err = db.QueryRow("select id, user_id, title, description, year from calendars where id = ?", lastID).Scan(&calendar.ID, &calendar.UserID, &calendar.Title, &calendar.Description, &calendar.Year)
	if err != nil {
		return nil, err
	}

	return &pb.Calendar{Id: calendar.ID, UserId: calendar.UserID, Title: calendar.Title, Description: calendar.Description, Year: calendar.Year}, nil
}

func main() {
	var err error
	db, err = sql.Open("mysql", "root@tcp(127.0.0.1:3306)/adventar_dev")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	log.Printf("listening at %s\n", port)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln(err)
	}

	s := grpc.NewServer()
	pb.RegisterAdventarServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
