package service

import (
	"context"
	"fmt"
	"strings"

	pb "github.com/adventar/adventar/api-server/grpc-server/grpc/adventar/v1"
	"golang.org/x/xerrors"
)

// ListCalendars lists calendars.
func (s *Service) ListCalendars(ctx context.Context, in *pb.ListCalendarsRequest) (*pb.ListCalendarsResponse, error) {
	conditionQueries := []string{"c.year = ?"}
	limitQuery := ""
	conditionValues := []interface{}{in.GetYear()}
	if in.GetUserId() != 0 {
		conditionQueries = append(conditionQueries, "c.user_id = ?")
		conditionValues = append(conditionValues, in.GetUserId())
	}
	if in.GetQuery() != "" {
		conditionQueries = append(conditionQueries, "(c.title like ? or c.description like ?)")
		conditionValues = append(conditionValues, "%"+in.GetQuery()+"%", "%"+in.GetQuery()+"%")
	}
	if in.GetPageSize() != 0 {
		limitQuery = "limit ?"
		conditionValues = append(conditionValues, in.GetPageSize())
	}
	sql := fmt.Sprintf(`
		select
			c.id,
			c.title,
			c.description,
			c.year,
			u.id,
			u.name,
			u.icon_url
		from calendars as c
		inner join users as u on u.id = c.user_id
		where %s
		order by c.id desc
		%s
	`, strings.Join(conditionQueries, " and "), limitQuery)

	rows, err := s.db.Query(sql, conditionValues...)
	if err != nil {
		return nil, xerrors.Errorf("Failed query to fetch calendars: %w", err)
	}
	defer rows.Close()
	var calendars []*pb.Calendar
	for rows.Next() {
		var calendar pb.Calendar
		var user pb.User
		err := rows.Scan(
			&calendar.Id,
			&calendar.Title,
			&calendar.Description,
			&calendar.Year,
			&user.Id,
			&user.Name,
			&user.IconUrl,
		)
		if err != nil {
			return nil, xerrors.Errorf("Failed to scan row: %w", err)
		}
		calendar.Owner = &user
		calendars = append(calendars, &calendar)
	}

	if len(calendars) != 0 {
		err := s.bindEntryCount(calendars)
		if err != nil {
			return nil, xerrors.Errorf("Failed to bind entry count: %w", err)
		}
	}

	return &pb.ListCalendarsResponse{Calendars: calendars}, nil
}
