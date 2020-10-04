package service

import (
	"context"
	"database/sql"

	pb "github.com/adventar/adventar/api-server/grpc-server/grpc/adventar/v1"
	"github.com/adventar/adventar/api-server/grpc-server/model"
	"github.com/adventar/adventar/api-server/grpc-server/util"
	"golang.org/x/xerrors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreateCalendar creates a calendar.
func (s *Service) CreateCalendar(ctx context.Context, in *pb.CreateCalendarRequest) (*pb.Calendar, error) {
	now, err := util.CurrentDate()
	if err != nil {
		return nil, err
	}

	if now.Month < 11 {
		return nil, status.Errorf(codes.FailedPrecondition, "Calendars can not create now.")
	}

	if in.GetTitle() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Title is invalid")
	}

	currentUser, err := s.getCurrentUser(ctx)
	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, "Authentication failed")
	}

	lastID, err := s.insertCalendar(currentUser.ID, in.GetTitle(), in.GetDescription(), now.Year)
	if err != nil {
		return nil, xerrors.Errorf("Failed to insert calendar: %w", err)
	}

	var calendar model.Calendar
	err = s.db.Get(&calendar, "select id, user_id, title, description from calendars where id = ?", lastID)
	if err == sql.ErrNoRows {
		return nil, status.Error(codes.NotFound, "Calendar not found")
	}
	if err != nil {
		return nil, xerrors.Errorf("Failed query to fetch calendar: %w", err)
	}

	return &pb.Calendar{Id: calendar.ID, Title: calendar.Title, Description: calendar.Description, Year: calendar.Year}, nil
}

func (s *Service) insertCalendar(userID int64, title string, description string, year int) (int64, error) {
	res, err := s.db.Exec(
		"insert into calendars(user_id, title, description, year) values(?, ?, ?, ?)",
		userID, title, description, year,
	)
	if err != nil {
		return 0, xerrors.Errorf("Failed query to insert into calendar: %w", err)
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return 0, xerrors.Errorf("Failed to get last id: %w", err)
	}

	return lastID, err
}
