package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	pb "github.com/adventar/adventar/grpc-server/adventar/v1"
)

type verifier interface {
	VerifyIDToken(string) *AuthResult
}

// Service holds data used by grpc functions.
type Service struct {
	db       *sql.DB
	verifier verifier
}

// NewService creates a new Service.
func NewService(db *sql.DB, verifier verifier) *Service {
	return &Service{db: db, verifier: verifier}
}

func (s *Service) serve(addr string) {
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

// ListCalendars lists calendars.
func (s *Service) ListCalendars(ctx context.Context, in *pb.ListCalendarsRequest) (*pb.ListCalendarsResponse, error) {
	return &pb.ListCalendarsResponse{}, nil
}

// GetCalendar returns a calendar.
func (s *Service) GetCalendar(ctx context.Context, in *pb.GetCalendarRequest) (*pb.GetCalendarResponse, error) {
	var calendar calendar
	row := s.db.QueryRow("select id, user_id, title, description, year from calendars where id = ?", in.GetCalendarId())
	err := row.Scan(&calendar.ID, &calendar.UserID, &calendar.Title, &calendar.Description, &calendar.Year)
	if err != nil {
		return nil, err
	}
	pbCalendar := &pb.Calendar{Id: calendar.ID, Title: calendar.Title, Description: calendar.Description, Year: calendar.Year}
	return &pb.GetCalendarResponse{Calendar: pbCalendar}, nil
}

// CreateCalendar creates a calendar.
func (s *Service) CreateCalendar(ctx context.Context, in *pb.CreateCalendarRequest) (*pb.Calendar, error) {
	currentUser, err := s.getCurrentUser(ctx)
	if err != nil {
		return nil, err
	}

	stmt, err := s.db.Prepare("insert into calendars(user_id, title, description, year) values(?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}

	res, err := stmt.Exec(currentUser.ID, in.GetTitle(), in.GetDescription(), time.Now().Year())
	if err != nil {
		return nil, err
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	var calendar calendar
	err = s.db.QueryRow("select id, user_id, title, description, year from calendars where id = ?", lastID).Scan(&calendar.ID, &calendar.UserID, &calendar.Title, &calendar.Description, &calendar.Year)
	if err != nil {
		return nil, err
	}

	return &pb.Calendar{Id: calendar.ID, Title: calendar.Title, Description: calendar.Description, Year: calendar.Year}, nil
}

// UpdateCalendar updates the calendar.
func (s *Service) UpdateCalendar(ctx context.Context, in *pb.UpdateCalendarRequest) (*pb.Calendar, error) {
	return &pb.Calendar{}, nil
}

// DeleteCalendar updates the calendar.
func (s *Service) DeleteCalendar(ctx context.Context, in *pb.DeleteCalendarRequest) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}

// CreateEntry creates a entry.
func (s *Service) CreateEntry(ctx context.Context, in *pb.CreateEntryRequest) (*pb.Entry, error) {
	return &pb.Entry{}, nil
}

// UpdateEntry updates the entry.
func (s *Service) UpdateEntry(ctx context.Context, in *pb.UpdateEntryRequest) (*pb.Entry, error) {
	return &pb.Entry{}, nil
}

// DeleteEntry deletes the entry.
func (s *Service) DeleteEntry(ctx context.Context, in *pb.DeleteEntryRequest) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}

// SignIn validates the id token.
func (s *Service) SignIn(ctx context.Context, in *pb.SignInRequest) (*empty.Empty, error) {
	authResult := s.verifier.VerifyIDToken(in.GetJwt())
	var userID int
	err := s.db.QueryRow("select id from users where auth_provider = ? and auth_uid = ?", authResult.AuthProvider, authResult.AuthUID).Scan(&userID)

	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if err == sql.ErrNoRows {
		stmt, err := s.db.Prepare("insert into users (name, auth_uid, auth_provider, icon_url) values (?, ?, ?, ?)")
		if err != nil {
			return nil, err
		}
		defer stmt.Close()
		_, err = stmt.Exec(authResult.Name, authResult.AuthUID, authResult.AuthProvider, authResult.IconURL)
		if err != nil {
			return nil, err
		}
	} else {
		stmt, err := s.db.Prepare("update users set icon_url = ? where id = ?")
		if err != nil {
			return nil, err
		}
		defer stmt.Close()
		_, err = stmt.Exec(authResult.IconURL, userID)
		if err != nil {
			return nil, err
		}
	}

	return &empty.Empty{}, nil
}

func (s *Service) getCurrentUser(ctx context.Context) (*user, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("not found metadata")
	}

	values := md["authorization"]
	if len(values) == 0 {
		return nil, fmt.Errorf("not found authorization in metadata")
	}

	authResult := s.verifier.VerifyIDToken(values[0])

	var user user
	err := s.db.QueryRow("select id, name, icon_url from users where auth_provider = ? and auth_uid = ?", authResult.AuthProvider, authResult.AuthUID).Scan(&user.ID, &user.Name, &user.IconURL)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
