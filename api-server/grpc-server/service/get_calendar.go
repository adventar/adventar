package service

import (
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	pb "github.com/adventar/adventar/api-server/grpc-server/grpc/adventar/v1"
	"github.com/adventar/adventar/api-server/grpc-server/model"
	"golang.org/x/xerrors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetCalendar returns a calendar.
func (s *Service) GetCalendar(ctx context.Context, in *pb.GetCalendarRequest) (*pb.GetCalendarResponse, error) {
	query, args, err := sq.
		Select(makeSelectValue(map[string][]string{
			"calendars": {"id", "title", "description", "year"},
			"users":     {"id", "name", "icon_url"},
		})...).
		From("calendars").
		Join("users on users.id = calendars.user_id").
		Where(sq.Eq{"calendars.id": in.GetCalendarId()}).
		ToSql()

	if err != nil {
		return nil, xerrors.Errorf("Failed query to create sql: %w", err)
	}

	result := struct {
		Calendar model.Calendar `db:"calendars"`
		User     model.User     `db:"users"`
	}{}

	err = s.db.Get(&result, query, args...)

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
