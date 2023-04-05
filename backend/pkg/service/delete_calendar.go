package service

import (
	"context"

	adventarv1 "github.com/adventar/adventar/backend/pkg/gen/proto/adventar/v1"
	"github.com/adventar/adventar/backend/pkg/usecase"
	"github.com/bufbuild/connect-go"
	"google.golang.org/protobuf/types/known/emptypb"
)

// DeleteCalendar deletes the calendar.
func (x *Service) DeleteCalendar(
	ctx context.Context,
	req *connect.Request[adventarv1.DeleteCalendarRequest],
) (*connect.Response[emptypb.Empty], error) {
	currentUser, err := x.authenticate(ctx)

	if err != nil {
		return nil, err
	}

	err = x.usecase.DeleteCalendar(&usecase.DeleteCalendarInput{
		CalendarID: req.Msg.GetCalendarId(),
		UserID:     currentUser.ID,
	})

	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&emptypb.Empty{}), nil
}
