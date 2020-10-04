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
	var result struct {
		Calendar model.Calendar `db:"c"`
		User     model.User     `db:"u"`
	}
	selectSQL := `
		select
			c.id as "c.id",
			c.title as "c.title",
			c.description as "c.description",
			c.year as "c.year",
			u.id as "u.id",
			u.name as "u.name",
			u.icon_url as "u.icon_url"
		from calendars as c
		inner join users as u on u.id = c.user_id
		where c.id = ?
	`

	err := s.db.Get(&result, selectSQL, in.GetCalendarId())

	if err == sql.ErrNoRows {
		return nil, status.Error(codes.NotFound, "Calendar not found")
	}

	if err != nil {
		return nil, xerrors.Errorf("Failed query to fetch calendar: %w", err)
	}

	calendar := result.Calendar
	user := result.User
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
