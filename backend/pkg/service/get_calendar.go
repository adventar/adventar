package service

import (
	"context"
	"database/sql"
	"errors"

	sq "github.com/Masterminds/squirrel"
	adventarv1 "github.com/adventar/adventar/backend/pkg/gen/adventar/v1"
	"github.com/adventar/adventar/backend/pkg/model"
	"github.com/bufbuild/connect-go"
	"github.com/m-mizutani/goerr"
	"golang.org/x/xerrors"
)

// GetCalendar returns a calendar.
func (s *Service) GetCalendar(
	ctx context.Context,
	req *connect.Request[adventarv1.GetCalendarRequest],
) (*connect.Response[adventarv1.GetCalendarResponse], error) {
	calendarId := req.Msg.GetCalendarId()
	query, args, err := sq.
		Select(makeSelectValue(map[string][]string{
			"calendars": {"id", "title", "description", "year"},
			"users":     {"id", "name", "icon_url"},
		})...).
		From("calendars").
		Join("users on users.id = calendars.user_id").
		Where(sq.Eq{"calendars.id": calendarId}).
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
		return nil, connect.NewError(connect.CodeNotFound, errors.New("Calendar not found"))
	}

	if err != nil {
		return nil, goerr.Wrap(err, "Failed query to fetch calendar").With("calendar_id", calendarId)
		// return nil, xerrors.Errorf("Failed query to fetch calendar: %w", err)
	}

	calendar := result.Calendar
	user := result.User
	entries, err := s.findEntries(calendar.ID)
	if err != nil {
		return nil, xerrors.Errorf("Failed to find entries: %w", err)
	}

	pbUser := &adventarv1.User{Id: user.ID, Name: user.Name, IconUrl: user.IconURL}
	pbCalendar := &adventarv1.Calendar{
		Id:          calendar.ID,
		Title:       calendar.Title,
		Description: calendar.Description,
		Year:        calendar.Year,
		Owner:       pbUser,
		EntryCount:  int32(len(entries)),
	}

	return connect.NewResponse(&adventarv1.GetCalendarResponse{
		Calendar: pbCalendar,
		Entries:  entries,
	}), nil
}
