package service

import (
	"context"
	"errors"

	adventarv1 "github.com/adventar/adventar/backend/pkg/gen/proto/adventar/v1"
	"github.com/bufbuild/connect-go"
	"github.com/m-mizutani/goerr"
	"google.golang.org/protobuf/types/known/emptypb"
)

// DeleteCalendar deletes the calendar.
func (s *Service) DeleteCalendar(
	ctx context.Context,
	req *connect.Request[adventarv1.DeleteCalendarRequest],
) (*connect.Response[emptypb.Empty], error) {
	currentUser, err := s.getCurrentUser(req.Header())
	if err != nil {
		return nil, connect.NewError(connect.CodePermissionDenied, errors.New("Authentication failed"))
	}

	_, err = s.db.Exec("delete from calendars where id = ? and user_id = ?", req.Msg.GetCalendarId(), currentUser.ID)
	if err != nil {
		return nil, goerr.Wrap(err, "Failed query to delete calendar")
	}

	return connect.NewResponse(&emptypb.Empty{}), nil
}
