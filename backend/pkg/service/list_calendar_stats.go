package service

import (
	"context"

	adventarv1 "github.com/adventar/adventar/backend/pkg/gen/adventar/v1"
	"github.com/bufbuild/connect-go"
	"github.com/m-mizutani/goerr"
)

// ListCalendarStats lists calendar stats
func (s *Service) ListCalendarStats(
	ctx context.Context,
	req *connect.Request[adventarv1.ListCalendarStatsRequest],
) (*connect.Response[adventarv1.ListCalendarStatsResponse], error) {
	type result struct {
		Year           int32 `db:"year"`
		CalendarsCount int32 `db:"calendars_count"`
		EntriesCount   int32 `db:"entries_count"`
	}

	sql := `
		select
			calendars.year as year
			, count(distinct calendars.id) as calendars_count
			, count(entries.id) as entries_count
		from
			calendars
			left join entries on entries.calendar_id = calendars.id
		group by
			year
		order by
			year desc
	`

	rows := []result{}
	err := s.db.Select(&rows, sql)
	if err != nil {
		return nil, goerr.Wrap(err, "Failed query to fetch calendar stats")
	}
	var stats []*adventarv1.CalendarStat
	for _, row := range rows {
		stat := adventarv1.CalendarStat{Year: row.Year, CalendarsCount: row.CalendarsCount, EntriesCount: row.EntriesCount}
		stats = append(stats, &stat)
	}

	return connect.NewResponse(&adventarv1.ListCalendarStatsResponse{CalendarStats: stats}), nil
}
