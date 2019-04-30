package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	pb "github.com/adventar/adventar/grpc-server/adventar/v1"
)

type Verifier interface {
	VerifyIDToken(string) *AuthResult
}

type Service struct {
	db       *sql.DB
	verifier Verifier
}

func NewService(db *sql.DB, verifier Verifier) *Service {
	return &Service{db: db, verifier: verifier}
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

func (s *Service) ListCalendars(ctx context.Context, in *pb.ListCalendarsRequest) (*pb.ListCalendarsResponse, error) {
	return &pb.ListCalendarsResponse{}, nil
}

func (s *Service) GetCalendar(ctx context.Context, in *pb.GetCalendarRequest) (*pb.GetCalendarResponse, error) {
	var calendar Calendar
	log.Printf("Request Id: %d", in.GetCalendarId())
	row := s.db.QueryRow("select id, user_id, title, description, year from calendars where id = ?", in.GetCalendarId())
	err := row.Scan(&calendar.ID, &calendar.UserID, &calendar.Title, &calendar.Description, &calendar.Year)
	if err != nil {
		return nil, err
	}
	pbCalendar := &pb.Calendar{Id: calendar.ID, Title: calendar.Title, Description: calendar.Description, Year: calendar.Year}
	return &pb.GetCalendarResponse{Calendar: pbCalendar}, nil
}

func (s *Service) CreateCalendar(ctx context.Context, in *pb.CreateCalendarRequest) (*pb.Calendar, error) {
	currentUser, err := s.GetCurrentUser(ctx)
	if err != nil {
		return nil, err
	}

	stmt, err := s.db.Prepare("insert into calendars(user_id, title, description, year) values(?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}

	res, err := stmt.Exec(currentUser.ID, in.GetTitle(), in.GetDescription(), in.GetYear())
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

	return &pb.Calendar{Id: calendar.ID, Title: calendar.Title, Description: calendar.Description, Year: calendar.Year}, nil
}

func (s *Service) UpdateCalendar(ctx context.Context, in *pb.UpdateCalendarRequest) (*pb.Calendar, error) {
	return &pb.Calendar{}, nil
}

func (s *Service) DeleteCalendar(ctx context.Context, in *pb.DeleteCalendarRequest) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}

func (s *Service) CreateEntry(ctx context.Context, in *pb.CreateEntryRequest) (*pb.Entry, error) {
	return &pb.Entry{}, nil
}

func (s *Service) UpdateEntry(ctx context.Context, in *pb.UpdateEntryRequest) (*pb.Entry, error) {
	return &pb.Entry{}, nil
}

func (s *Service) DeleteEntry(ctx context.Context, in *pb.DeleteEntryRequest) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}

func (s *Service) SignIn(ctx context.Context, in *pb.SignInRequest) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}

func (s *Service) GetCurrentUser(ctx context.Context) (*User, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("not found metadata")
	}

	values := md["authorization"]
	if len(values) == 0 {
		return nil, fmt.Errorf("not found authorization in metadata")
	}

	authResult := s.verifier.VerifyIDToken(values[0])

	var user User
	err := s.db.QueryRow("select id, name, icon_url from users where auth_provider = ? and auth_uid = ?", authResult.AuthProvider, authResult.AuthUID).Scan(&user.ID, &user.Name, &user.IconURL)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
