package controller

import (
	"context"

	"github.com/adventar/adventar/backend/pkg/domain/model"
	adventarv1 "github.com/adventar/adventar/backend/pkg/gen/proto/adventar/v1"
	"github.com/bufbuild/connect-go"
	"github.com/m-mizutani/goerr"
	"github.com/m-mizutani/gots/slice"
)

// ListCalendarStats lists calendar stats
func (x *Controller) ListCalendarStats(
	ctx context.Context,
	req *connect.Request[adventarv1.ListCalendarStatsRequest],
) (*connect.Response[adventarv1.ListCalendarStatsResponse], error) {
	stats, err := x.usecase.ListCalendarStats()

	if err != nil {
		return nil, goerr.Wrap(err, "Failed to list calendar stats")
	}

	return connect.NewResponse(&adventarv1.ListCalendarStatsResponse{
		CalendarStats: slice.Map(stats, func(stat *model.CalendarStat) *adventarv1.CalendarStat {
			return stat.ToProto()
		}),
	}), nil
}
