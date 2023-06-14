package controller

import (
	"context"

	adventarv1 "github.com/adventar/adventar/backend/pkg/gen/proto/adventar/v1"
	"github.com/adventar/adventar/backend/pkg/usecase"
	"github.com/bufbuild/connect-go"
	"github.com/m-mizutani/goerr"
)

// UpdateCalendar updates the calendar.
func (x *Service) UpdateCalendar(
	ctx context.Context,
	req *connect.Request[adventarv1.UpdateCalendarRequest],
) (*connect.Response[adventarv1.Calendar], error) {
	currentUser, err := x.authenticate(ctx)

	if err != nil {
		return nil, goerr.Wrap(err, "Failed to authenticate")
	}

	calendar, err := x.usecase.UpdateCalendar(&usecase.UpdateCalendarInput{
		CalendarID:  req.Msg.GetCalendarId(),
		Title:       req.Msg.GetTitle(),
		Description: req.Msg.GetDescription(),
		UserID:      currentUser.ID,
	})

	return connect.NewResponse(calendar.ToProto()), nil
}
