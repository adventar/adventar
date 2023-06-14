package controller

import (
	"context"

	adventarv1 "github.com/adventar/adventar/backend/pkg/gen/proto/adventar/v1"
	"github.com/adventar/adventar/backend/pkg/usecase"
	"github.com/bufbuild/connect-go"
	"github.com/m-mizutani/goerr"
	"google.golang.org/protobuf/types/known/emptypb"
)

// DeleteCalendar deletes the calendar.
func (x *Controller) DeleteCalendar(
	ctx context.Context,
	req *connect.Request[adventarv1.DeleteCalendarRequest],
) (*connect.Response[emptypb.Empty], error) {
	currentUser, err := x.authenticate(ctx)

	if err != nil {
		return nil, goerr.Wrap(err, "Failed to authenticate")
	}

	err = x.usecase.DeleteCalendar(&usecase.DeleteCalendarInput{
		CalendarID: req.Msg.GetCalendarId(),
		UserID:     currentUser.ID,
	})

	if err != nil {
		return nil, goerr.Wrap(err, "Failed to delete calendar")
	}

	return connect.NewResponse(&emptypb.Empty{}), nil
}
