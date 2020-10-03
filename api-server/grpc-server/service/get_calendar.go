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

// GetCalendar returns a calendar.
func (s *Service) GetCalendar(ctx context.Context, in *pb.GetCalendarRequest) (*pb.GetCalendarResponse, error) {
	var calendar model.Calendar
	var user model.User
	selectSQL := `
		select
			c.id,
			c.title,
			c.description,
			c.year,
			u.id,
			u.name,
			u.icon_url
		from calendars as c
		inner join users as u on u.id = c.user_id
		where c.id = ?
	`

	row := s.db.QueryRow(selectSQL, in.GetCalendarId())
	err := row.Scan(
		&calendar.ID,
		&calendar.Title,
		&calendar.Description,
		&calendar.Year,
		&user.ID,
		&user.Name,
		&user.IconURL,
	)

	if err == sql.ErrNoRows {
		return nil, status.Error(codes.NotFound, "Calendar not found")
	}

	if err != nil {
		return nil, xerrors.Errorf("Failed query to fetch calendar: %w", err)
	}

	entries, err := s.findEntries(calendar.ID)
	if err != nil {
		return nil, xerrors.Errorf("Failed to find entries: %w", err)
	}

	pbUser := &pb.User{Id: user.ID, Name: user.Name, IconUrl: user.IconURL}
	pbCalendar := &pb.Calendar{
		Id:          calendar.ID,
		Title:       calendar.Title,
		Description: calendar.Description,
		Year:        calendar.Year,
		Owner:       pbUser,
		EntryCount:  int32(len(entries)),
	}

	return &pb.GetCalendarResponse{Calendar: pbCalendar, Entries: entries}, nil
}
