package service

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"golang.org/x/xerrors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/adventar/adventar/backend/grpc-server/grpc/adventar/v1"
	"github.com/golang/protobuf/ptypes/empty"
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
		entries = append(entries, &e)
	}

	return &pb.ListEntriesResponse{Entries: entries}, nil
}

// CreateEntry creates a entry.
func (s *Service) CreateEntry(ctx context.Context, in *pb.CreateEntryRequest) (*pb.Entry, error) {
	currentUser, err := s.getCurrentUser(ctx)
	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, "Invalid token")
	}

	var year int
	err = s.db.QueryRow("select year from calendars where id = ?", in.GetCalendarId()).Scan(&year)
	if err == sql.ErrNoRows {
		return nil, status.Errorf(codes.NotFound, "Calendar not found")
	}
	if err != nil {
		return nil, xerrors.Errorf("Failed query to fetch calendar: %w", err)
	}

	day := in.GetDay()
	if day < 1 || day > 25 {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid day: %d", day)
	}

	stmt, err := s.db.Prepare("insert into entries(user_id, calendar_id, day, comment, url, title, image_url) values(?, ?, ?, '', '', '', '')")
	if err != nil {
		return nil, xerrors.Errorf("Failed to prepare query: %w", err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(currentUser.ID, in.GetCalendarId(), day)
	if err != nil {
		return nil, xerrors.Errorf("Failed query to insert into entry: %w", err)
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return nil, xerrors.Errorf("Failed to get last id: %w", err)
	}

	var entryID int64
	err = s.db.QueryRow("select id from entries where id = ?", lastID).Scan(&entryID)
	if err != nil {
		return nil, xerrors.Errorf("Failed query to fetch entry: %w", err)
	}

	return &pb.Entry{Id: entryID}, nil
}

// UpdateEntry updates the entry.
func (s *Service) UpdateEntry(ctx context.Context, in *pb.UpdateEntryRequest) (*pb.Entry, error) {
	currentUser, err := s.getCurrentUser(ctx)
	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, "Invalid token")
	}

	stmt, err := s.db.Prepare("update entries set comment = ?, url = ? where id = ? and user_id = ?")
	if err != nil {
		return nil, xerrors.Errorf("Failed to prepare query: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(in.GetComment(), in.GetUrl(), in.GetEntryId(), currentUser.ID)
	if err != nil {
		return nil, xerrors.Errorf("Failed query to update entry: %w", err)
	}

	if in.GetUrl() != "" {
		m, err := s.metaFetcher.Fetch(in.GetUrl())
		// TODO: Ignore error
		if err != nil {
			return nil, xerrors.Errorf("Failed to fetch url: %w", err)
		}
		stmt, err = s.db.Prepare("update entries set title = ?, image_url = ? where id = ? and user_id = ?")
		if err != nil {
			return nil, xerrors.Errorf("Failed to prepare query: %w", err)
		}
		defer stmt.Close()

		_, err = stmt.Exec(m.Title, m.ImageURL, in.GetEntryId(), currentUser.ID)
		if err != nil {
			return nil, xerrors.Errorf("Failed query to update entry: %w", err)
		}
	}

	var comment string
	var url string
	var title string
	var imageURL string
	err = s.db.QueryRow("select comment, url, title, image_url from entries where id = ?", in.GetEntryId()).Scan(&comment, &url, &title, &imageURL)
	if err == sql.ErrNoRows {
		return nil, status.Errorf(codes.NotFound, "Entry not found")
	}
	if err != nil {
		return nil, xerrors.Errorf("Failed query to fetch entry: %w", err)
	}

	return &pb.Entry{Id: in.GetEntryId(), Comment: comment, Url: url, Title: title, ImageUrl: imageURL}, nil
}

// DeleteEntry deletes the entry.
func (s *Service) DeleteEntry(ctx context.Context, in *pb.DeleteEntryRequest) (*empty.Empty, error) {
	currentUser, err := s.getCurrentUser(ctx)
	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, "Invalid token")
	}

	deletable, err := s.entryDeletable(int(in.GetEntryId()), int(currentUser.ID))
	if err != nil {
		return nil, err
	}
	if deletable == false {
		return nil, status.Errorf(codes.PermissionDenied, "Invalid request")
	}

	stmt, err := s.db.Prepare("delete from entries where id = ?")
	if err != nil {
		return nil, xerrors.Errorf("Failed to prepare entry: %w", err)
	}

	_, err = stmt.Exec(in.GetEntryId())
	if err != nil {
		return nil, xerrors.Errorf("Failed query to delete entry: %w", err)
	}

	return &empty.Empty{}, nil
}

func (s *Service) entryDeletable(entryID int, userID int) (bool, error) {
	var entryOwnerID int
	var day int
	var year int
	var calendarOwnerID int
	err := s.db.QueryRow(`
		select
			e.user_id,
			e.day,
			c.year,
			c.user_id
		from
			entries as e
			inner join calendars as c on c.id = e.calendar_id
		where
			e.id = ?
	`, entryID).Scan(
		&entryOwnerID,
		&day,
		&year,
		&calendarOwnerID,
	)
	if err != nil {
		return false, xerrors.Errorf("Failed query to fetch user: %w", err)
	}

	if userID == entryOwnerID {
		return true, nil
	}

	now, err := currentDate()
	if err != nil {
		return false, err
	}

	if userID == calendarOwnerID && (year < now.Year || day+1 < now.Day) {
		return true, nil
	}

	return false, nil
}
