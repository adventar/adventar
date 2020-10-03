package service

import (
	"context"
	"net/url"

	pb "github.com/adventar/adventar/api-server/grpc-server/grpc/adventar/v1"
	"github.com/adventar/adventar/api-server/grpc-server/util"
	"github.com/golang/protobuf/ptypes/empty"
	"golang.org/x/xerrors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// DeleteEntry deletes the entry.
func (s *Service) DeleteEntry(ctx context.Context, in *pb.DeleteEntryRequest) (*empty.Empty, error) {
	currentUser, err := s.getCurrentUser(ctx)
	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, "Authentication failed")
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

	now, err := util.CurrentDate()
	if err != nil {
		return false, err
	}

	if userID == calendarOwnerID && (year < now.Year || day+1 < now.Day) {
		return true, nil
	}

	return false, nil
}

func isValidURL(s string) bool {
	u, err := url.Parse(s)
	if err != nil {
		return false
	}

	return u.Scheme == "http" || u.Scheme == "https"
}
