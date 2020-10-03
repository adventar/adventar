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

	stmt, err := s.db.Prepare("update calendars set title = ?, description = ? where id = ? and user_id = ?")
	if err != nil {
		return nil, xerrors.Errorf("Failed to prepare query: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(in.GetTitle(), in.GetDescription(), in.GetCalendarId(), currentUser.ID)
	if err != nil {
		return nil, xerrors.Errorf("Failed query to update calendar: %w", err)
	}

	var calendar model.Calendar
	err = s.db.QueryRow("select id, user_id, title, description, year from calendars where id = ? and user_id = ?", in.GetCalendarId(), currentUser.ID).Scan(&calendar.ID, &calendar.UserID, &calendar.Title, &calendar.Description, &calendar.Year)
	if err == sql.ErrNoRows {
		return nil, status.Error(codes.NotFound, "Calendar not found")
	}
	if err != nil {
		return nil, xerrors.Errorf("Failed query to fetch calendar: %w", err)
	}

	return &pb.Calendar{Id: calendar.ID, Title: calendar.Title, Description: calendar.Description, Year: calendar.Year}, nil
}
