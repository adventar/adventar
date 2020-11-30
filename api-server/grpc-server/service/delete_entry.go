package service

import (
	"context"

	pb "github.com/adventar/adventar/api-server/grpc-server/grpc/adventar/v1"
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

	_, err = s.db.Exec("delete from entries where id = ?", in.GetEntryId())
	if err != nil {
		return nil, xerrors.Errorf("Failed to delete entry: %w", err)
	}

	return &empty.Empty{}, nil
}

func (s *Service) entryDeletable(entryID int, userID int) (bool, error) {
	var result struct {
		EntryOwnerID    int `db:"entry_owner_id"`
		CalendarOwnerID int `db:"calendar_owner_id"`
	}

	sql := `
		select
			e.user_id as entry_owner_id,
			c.user_id as calendar_owner_id
		from
			entries as e
			inner join calendars as c on c.id = e.calendar_id
		where
			e.id = ?
	`

	err := s.db.Get(&result, sql, entryID)

	if err != nil {
		return false, xerrors.Errorf("Failed query to fetch user: %w", err)
	}

	if userID == result.EntryOwnerID {
		return true, nil
	}

	if userID == result.CalendarOwnerID {
		return true, nil
	}

	return false, nil
}
