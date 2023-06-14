package controller

import (
	"context"

	adventarv1 "github.com/adventar/adventar/backend/pkg/gen/proto/adventar/v1"
	"github.com/adventar/adventar/backend/pkg/usecase"
	"github.com/bufbuild/connect-go"
	"github.com/m-mizutani/goerr"
)

// CreateCalendar creates a calendar.
func (x *Controller) CreateCalendar(
	ctx context.Context,
	req *connect.Request[adventarv1.CreateCalendarRequest],
) (*connect.Response[adventarv1.Calendar], error) {
	currentUser, err := x.authenticate(ctx)

	if err != nil {
		return nil, goerr.Wrap(err, "Failed to authenticate")
	}

	calendar, err := x.usecase.CreateCalendar(&usecase.CreateCalendarInput{
		Title:       req.Msg.GetTitle(),
		Description: req.Msg.GetDescription(),
		UserID:      currentUser.ID,
	})

	if err != nil {
		return nil, goerr.Wrap(err, "Failed to create calendar")
	}

	return connect.NewResponse(calendar.ToProto()), nil
}
