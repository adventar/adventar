package service

import (
	"context"
	"fmt"
	"strings"

	pb "github.com/adventar/adventar/api-server/grpc-server/grpc/adventar/v1"
	"golang.org/x/xerrors"
)

// ListEntries lists entries.
func (s *Service) ListEntries(ctx context.Context, in *pb.ListEntriesRequest) (*pb.ListEntriesResponse, error) {
	conditionQueries := []string{"e.user_id = ?"}
	conditionValues := []interface{}{in.GetUserId()}

	if in.GetYear() != 0 {
		conditionQueries = append(conditionQueries, "c.year = ?")
		conditionValues = append(conditionValues, in.GetYear())
	}

	sql := fmt.Sprintf(`
		select
			e.id,
			e.day,
			e.title,
			e.comment,
			e.url,
			e.image_url,
			c.id,
			c.title,
			c.description,
			c.year,
			u.id,
			u.name,
			u.icon_url
		from entries as e
		inner join users as u on u.id = e.user_id
		inner join calendars as c on c.id = e.calendar_id
		where %s
		order by c.year, e.day, e.id
	`, strings.Join(conditionQueries, " and "))

	rows, err := s.db.Query(sql, conditionValues...)
	defer rows.Close()
	if err != nil {
		return nil, xerrors.Errorf("Failed query to fetch entries: %w", err)
	}

	entries := []*pb.Entry{}
	for rows.Next() {
		var e pb.Entry
		var c pb.Calendar
		var u pb.User
		err := rows.Scan(
			&e.Id,
			&e.Day,
			&e.Title,
			&e.Comment,
			&e.Url,
			&e.ImageUrl,
			&c.Id,
			&c.Title,
			&c.Description,
			&c.Year,
			&u.Id,
			&u.Name,
			&u.IconUrl,
		)
		if err != nil {
			return nil, xerrors.Errorf("Failed to scan row: %w", err)
		}
		e.Calendar = &c
		e.Owner = &u
		e.ImageUrl = convertImageURL(e.ImageUrl)
		entries = append(entries, &e)
	}

	return &pb.ListEntriesResponse{Entries: entries}, nil
}
