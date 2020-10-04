package service

import (
	"context"

	pb "github.com/adventar/adventar/api-server/grpc-server/grpc/adventar/v1"
	"golang.org/x/xerrors"
)

// ListCalendarStats lists calendar stats
func (s *Service) ListCalendarStats(ctx context.Context, in *pb.ListCalendarStatsRequest) (*pb.ListCalendarStatsResponse, error) {
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
		return nil, xerrors.Errorf("Failed query to fetch calendar stats: %w", err)
	}
	var stats []*pb.CalendarStat
	for _, row := range rows {
		stat := pb.CalendarStat{Year: row.Year, CalendarsCount: row.CalendarsCount, EntriesCount: row.EntriesCount}
		stats = append(stats, &stat)
	}

	return &pb.ListCalendarStatsResponse{CalendarStats: stats}, nil
}
