package service

import (
	"context"
	"database/sql"

	pb "github.com/adventar/adventar/api-server/grpc-server/grpc/adventar/v1"
	"github.com/adventar/adventar/api-server/grpc-server/model"
	"golang.org/x/xerrors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// UpdateCalendar updates the calendar.
func (s *Service) UpdateCalendar(ctx context.Context, in *pb.UpdateCalendarRequest) (*pb.Calendar, error) {
	currentUser, err := s.getCurrentUser(ctx)
	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, "Authentication failed")
	}
	if in.GetTitle() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Title is invalid")
	}

	_, err = s.db.Exec(
		"update calendars set title = ?, description = ? where id = ? and user_id = ?",
		in.GetTitle(), in.GetDescription(), in.GetCalendarId(), currentUser.ID,
	)
	if err != nil {
		return nil, xerrors.Errorf("Failed query to update calendar: %w", err)
	}

	var calendar model.Calendar
	err = s.db.Get(&calendar, "select * from calendars where id = ? and user_id = ?", in.GetCalendarId(), currentUser.ID)
	if err == sql.ErrNoRows {
		return nil, status.Error(codes.NotFound, "Calendar not found")
	}
	if err != nil {
		return nil, xerrors.Errorf("Failed query to fetch calendar: %w", err)
	}

	return &pb.Calendar{Id: calendar.ID, Title: calendar.Title, Description: calendar.Description, Year: calendar.Year}, nil
}
