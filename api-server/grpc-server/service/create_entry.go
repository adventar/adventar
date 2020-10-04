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
	err = s.db.Get(&year, "select year from calendars where id = ?", in.GetCalendarId())
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

	lastID, err := s.insertEntry(currentUser.ID, in.GetCalendarId(), day)
	if err != nil {
		return nil, xerrors.Errorf("Failed to insert entry: %w", err)
	}

	var entryID int64
	err = s.db.Get(&entryID, "select id from entries where id = ?", lastID)
	if err != nil {
		return nil, xerrors.Errorf("Failed query to fetch entry: %w", err)
	}

	return &pb.Entry{Id: entryID}, nil
}

func (s *Service) insertEntry(userID int64, calendarID int64, day int32) (int64, error) {
	res, err := s.db.Exec(
		"insert into entries(user_id, calendar_id, day, comment, url, title, image_url) values(?, ?, ?, '', '', '', '')",
		userID, calendarID, day,
	)
	if err != nil {
		return 0, xerrors.Errorf("Failed query to insert into entry: %w", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, xerrors.Errorf("Failed to get last id: %w", err)
	}

	return id, nil
}
