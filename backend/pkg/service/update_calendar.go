package service

import (
	"context"
	"database/sql"
	"errors"

	adventarv1 "github.com/adventar/adventar/backend/pkg/gen/adventar/v1"
	"github.com/adventar/adventar/backend/pkg/model"
	"github.com/bufbuild/connect-go"
	"github.com/m-mizutani/goerr"
)

// UpdateCalendar updates the calendar.
func (s *Service) UpdateCalendar(
	ctx context.Context,
	req *connect.Request[adventarv1.UpdateCalendarRequest],
) (*connect.Response[adventarv1.Calendar], error) {
	currentUser, err := s.getCurrentUser(req.Header())
	if err != nil {
		return nil, connect.NewError(connect.CodePermissionDenied, errors.New("Authentication failed"))
	}
	if req.Msg.GetTitle() == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("Title is invalid"))
	}

	_, err = s.db.Exec(
		"update calendars set title = ?, description = ? where id = ? and user_id = ?",
		req.Msg.GetTitle(), req.Msg.GetDescription(), req.Msg.GetCalendarId(), currentUser.ID,
	)
	if err != nil {
		return nil, goerr.Wrap(err, "Failed query to update calendar")
	}

	var calendar model.Calendar
	err = s.db.Get(&calendar, "select id, user_id, title, description, year from calendars where id = ? and user_id = ?", req.Msg.GetCalendarId(), currentUser.ID)
	if err == sql.ErrNoRows {
		return nil, connect.NewError(connect.CodeNotFound, errors.New("Calendar not found"))
	}
	if err != nil {
		return nil, goerr.Wrap(err, "Failed query to fetch calendar")
	}

	return connect.NewResponse(&adventarv1.Calendar{
		Id:          calendar.ID,
		Title:       calendar.Title,
		Description: calendar.Description,
		Year:        calendar.Year,
	}), nil
}
