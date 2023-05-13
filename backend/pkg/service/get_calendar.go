package service

import (
	"context"

	"github.com/adventar/adventar/backend/pkg/domain/model"
	adventarv1 "github.com/adventar/adventar/backend/pkg/gen/proto/adventar/v1"
	"github.com/bufbuild/connect-go"
	"github.com/m-mizutani/goerr"
	"github.com/m-mizutani/gots/slice"
)

// GetCalendar returns a calendar.
func (x *Service) GetCalendar(
	ctx context.Context,
	req *connect.Request[adventarv1.GetCalendarRequest],
) (*connect.Response[adventarv1.GetCalendarResponse], error) {
	calendarId := req.Msg.GetCalendarId()
	calendar, err := x.usecase.GetCalendarById(calendarId)

	if err != nil {
		return nil, goerr.Wrap(err, "Failed to get calendar")
	}

	return connect.NewResponse(&adventarv1.GetCalendarResponse{
		Calendar: calendar.ToProto(),
		Entries: slice.Map(calendar.Entries, func(entry *model.Entry) *adventarv1.Entry {
			return entry.ToProto()
		}),
	}), nil
}
