package service

import (
	"context"
	"errors"

	adventarv1 "github.com/adventar/adventar/backend/pkg/gen/adventar/v1"
	"github.com/bufbuild/connect-go"
	"github.com/golang/protobuf/ptypes/empty"
	"golang.org/x/xerrors"
)

// DeleteCalendar deletes the calendar.
func (s *Service) DeleteCalendar(
	ctx context.Context,
	req *connect.Request[adventarv1.DeleteCalendarRequest],
) (*connect.Response[empty.Empty], error) {
	currentUser, err := s.getCurrentUser(req.Header())
	if err != nil {
		return nil, connect.NewError(connect.CodePermissionDenied, errors.New("Authentication failed"))
	}

	_, err = s.db.Exec("delete from calendars where id = ? and user_id = ?", req.Msg.GetCalendarId(), currentUser.ID)
	if err != nil {
		return nil, xerrors.Errorf("Failed query to delete calendar: %w", err)
	}

	return connect.NewResponse(&empty.Empty{}), nil
}
