package main

import (
	"context"
	"database/sql"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/adventar/adventar/grpc-server/adventar/v1"
)

type Service struct {
	db *sql.DB
}

type Calendar struct {
	ID          int64
	UserID      int64
	Title       string
	Description string
	Year        int32
}

func NewService(db *sql.DB) *Service {
	return &Service{db: db}
}

func (s *Service) Serve(addr string) {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	pb.RegisterAdventarServer(server, s)
	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *Service) GetCalendar(ctx context.Context, in *pb.GetCalendarRequest) (*pb.Calendar, error) {
	var calendar Calendar
	log.Printf("Request Id: %d", in.GetId())
	row := s.db.QueryRow("select id, user_id, title, description, year from calendars where id = ?", in.GetId())
	err := row.Scan(&calendar.ID, &calendar.UserID, &calendar.Title, &calendar.Description, &calendar.Year)
	if err != nil {
		return nil, err
	}
	return &pb.Calendar{Id: calendar.ID, UserId: calendar.UserID, Title: calendar.Title, Description: calendar.Description, Year: calendar.Year}, nil
}

func (s *Service) CreateCalendar(ctx context.Context, in *pb.CreateCalendarRequest) (*pb.Calendar, error) {
	stmt, err := s.db.Prepare("insert into calendars(user_id, title, description, year) values(?, ?, ?, ?)")
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
	err = s.db.QueryRow("select id, user_id, title, description, year from calendars where id = ?", lastID).Scan(&calendar.ID, &calendar.UserID, &calendar.Title, &calendar.Description, &calendar.Year)
	if err != nil {
		return nil, err
	}

	return &pb.Calendar{Id: calendar.ID, UserId: calendar.UserID, Title: calendar.Title, Description: calendar.Description, Year: calendar.Year}, nil
}
