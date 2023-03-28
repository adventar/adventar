package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	adventarv1 "github.com/adventar/adventar/backend/pkg/gen/proto/adventar/v1"
	"github.com/bufbuild/connect-go"
	"github.com/m-mizutani/goerr"
)

// CreateEntry creates a entry.
func (s *Service) CreateEntry(
	ctx context.Context,
	req *connect.Request[adventarv1.CreateEntryRequest],
) (*connect.Response[adventarv1.Entry], error) {
	currentUser, err := s.getCurrentUser(req.Header())
	if err != nil {
		return nil, connect.NewError(connect.CodePermissionDenied, errors.New("Authentication failed"))
	}

	var year int
	err = s.db.Get(&year, "select year from calendars where id = ?", req.Msg.GetCalendarId())
	if err == sql.ErrNoRows {
		return nil, connect.NewError(connect.CodeNotFound, errors.New("Calendar not found"))
	}
	if err != nil {
		return nil, goerr.Wrap(err, "Failed query to fetch calendar")
	}

	day := req.Msg.GetDay()
	if day < 1 || day > 25 {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("Invalid day: %d", day))
	}

	lastID, err := s.insertEntry(currentUser.ID, req.Msg.GetCalendarId(), day)
	if err != nil {
		return nil, goerr.Wrap(err, "Failed to insert entry")
	}

	var entryID int64
	err = s.db.Get(&entryID, "select id from entries where id = ?", lastID)
	if err != nil {
		return nil, goerr.Wrap(err, "Failed query to fetch entry")
	}

	return connect.NewResponse(&adventarv1.Entry{Id: entryID}), nil
}

func (s *Service) insertEntry(userID int64, calendarID int64, day int32) (int64, error) {
	res, err := s.db.Exec(
		"insert into entries(user_id, calendar_id, day, comment, url, title, image_url) values(?, ?, ?, '', '', '', '')",
		userID, calendarID, day,
	)
	if err != nil {
		return 0, goerr.Wrap(err, "Failed query to insert into entry")
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, goerr.Wrap(err, "Failed to get last id")
	}

	return id, nil
}
