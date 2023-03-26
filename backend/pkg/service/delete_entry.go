package service

import (
	"context"
	"errors"

	adventarv1 "github.com/adventar/adventar/backend/pkg/gen/adventar/v1"
	"github.com/bufbuild/connect-go"
	"github.com/golang/protobuf/ptypes/empty"
	"golang.org/x/xerrors"
)

// DeleteEntry deletes the entry.
func (s *Service) DeleteEntry(
	ctx context.Context,
	req *connect.Request[adventarv1.DeleteEntryRequest],
) (*connect.Response[empty.Empty], error) {
	currentUser, err := s.getCurrentUser(req.Header())
	if err != nil {
		return nil, connect.NewError(connect.CodePermissionDenied, errors.New("Authentication failed"))
	}

	deletable, err := s.entryDeletable(int(req.Msg.GetEntryId()), int(currentUser.ID))
	if err != nil {
		return nil, err
	}
	if deletable == false {
		return nil, connect.NewError(connect.CodePermissionDenied, errors.New("Invalid request"))
	}

	_, err = s.db.Exec("delete from entries where id = ?", req.Msg.GetEntryId())
	if err != nil {
		return nil, xerrors.Errorf("Failed to delete entry: %w", err)
	}

	return connect.NewResponse(&empty.Empty{}), nil
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
