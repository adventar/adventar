package controller

import (
	"context"

	"github.com/adventar/adventar/backend/pkg/domain/model"
	adventarv1 "github.com/adventar/adventar/backend/pkg/gen/proto/adventar/v1"
	"github.com/bufbuild/connect-go"
	"github.com/m-mizutani/goerr"
	"github.com/m-mizutani/gots/slice"
)

// ListCalendars lists calendars.
func (x *Service) ListCalendars(
	ctx context.Context,
	req *connect.Request[adventarv1.ListCalendarsRequest],
) (*connect.Response[adventarv1.ListCalendarsResponse], error) {
	year := req.Msg.Year
	userID := req.Msg.UserId
	limit := req.Msg.PageSize
	query := req.Msg.Query

	var calendars []*model.Calendar
	var err error
	if userID != 0 {
		calendars, err = x.usecase.ListCalendarsByUserId(year, userID)
	} else if query != "" {
		calendars, err = x.usecase.SearchCalendars(year, query)
	} else if limit != 0 {
		calendars, err = x.usecase.ListCalendars(year, limit)
	} else {
		calendars, err = x.usecase.ListAllCalendars(year)
	}

	if err != nil {
		return nil, goerr.Wrap(err, "Failed to list calendars")
	}

	return connect.NewResponse(&adventarv1.ListCalendarsResponse{
		Calendars: slice.Map(calendars, func(calendar *model.Calendar) *adventarv1.Calendar {
			return calendar.ToProto()
		}),
	}), nil
}
