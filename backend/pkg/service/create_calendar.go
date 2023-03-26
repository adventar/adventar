package service

import (
	"context"
	"database/sql"
	"errors"

	adventarv1 "github.com/adventar/adventar/backend/pkg/gen/adventar/v1"
	"github.com/adventar/adventar/backend/pkg/model"
	"github.com/adventar/adventar/backend/pkg/util"
	"github.com/bufbuild/connect-go"
	"golang.org/x/xerrors"
)

// CreateCalendar creates a calendar.
func (s *Service) CreateCalendar(
	ctx context.Context,
	req *connect.Request[adventarv1.CreateCalendarRequest],
) (*connect.Response[adventarv1.Calendar], error) {
	now, err := util.CurrentDate()
	if err != nil {
		return nil, err
	}

	if now.Month < 11 {
		return nil, connect.NewError(connect.CodeFailedPrecondition, errors.New("Calendars can not create now"))
	}

	if req.Msg.GetTitle() == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("Title is invalid"))
	}

	currentUser, err := s.getCurrentUser(req.Header())
	if err != nil {
		return nil, connect.NewError(connect.CodePermissionDenied, errors.New("Authentication failed"))
	}

	lastID, err := s.insertCalendar(currentUser.ID, req.Msg.GetTitle(), req.Msg.GetDescription(), now.Year)
	if err != nil {
		return nil, xerrors.Errorf("Failed to insert calendar: %w", err)
	}

	var calendar model.Calendar
	err = s.db.Get(&calendar, "select id, user_id, title, description from calendars where id = ?", lastID)
	if err == sql.ErrNoRows {
		return nil, connect.NewError(connect.CodeNotFound, errors.New("Calendar not found"))
	}
	if err != nil {
		return nil, xerrors.Errorf("Failed query to fetch calendar: %w", err)
	}

	return connect.NewResponse(&adventarv1.Calendar{
		Id:          calendar.ID,
		Title:       calendar.Title,
		Description: calendar.Description,
		Year:        calendar.Year,
	}), nil
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
