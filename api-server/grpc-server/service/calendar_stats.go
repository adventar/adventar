package service

import (
	"context"

	"golang.org/x/xerrors"
	pb "github.com/adventar/adventar/api-server/grpc-server/grpc/adventar/v1"
)

// ListCalendarStats lists calendar stats
func (s *Service) ListCalendarStats(ctx context.Context, in *pb.ListCalendarStatsRequest) (*pb.ListCalendarStatsResponse, error) {
	sql := `
		select
			calendars.year
			, count(distinct calendars.id) calendars_count
			, count(entries.id) entries_count
		from
			calendars
			left join entries on entries.calendar_id = calendars.id
		group by
			year
		order by
			year desc
	`
	rows, err := s.db.Query(sql)
	if err != nil {
		return nil, xerrors.Errorf("Failed query to fetch calendar stats: %w", err)
	}
	defer rows.Close()
	var stats []*pb.CalendarStat
	for rows.Next() {
		var stat pb.CalendarStat
		err := rows.Scan(
			&stat.Year,
			&stat.CalendarsCount,
			&stat.EntriesCount,
		)
		if err != nil {
			return nil, xerrors.Errorf("Failed to scan row: %w", err)
		}
		stats = append(stats, &stat)
	}

	return &pb.ListCalendarStatsResponse{CalendarStats: stats}, nil
}
