package service

import (
	"context"
	"database/sql"

	pb "github.com/adventar/adventar/api-server/grpc-server/grpc/adventar/v1"
	"golang.org/x/xerrors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreateEntry creates a entry.
func (s *Service) CreateEntry(ctx context.Context, in *pb.CreateEntryRequest) (*pb.Entry, error) {
	currentUser, err := s.getCurrentUser(ctx)
	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, "Authentication failed")
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
