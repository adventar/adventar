package usecase

import (
	"context"
	"database/sql"

	"github.com/adventar/adventar/backend/pkg/domain/model"
	"github.com/adventar/adventar/backend/pkg/domain/types"
	"github.com/adventar/adventar/backend/pkg/gen/sqlc/adventar_db"
	"github.com/m-mizutani/goerr"
)

func (x *Usecase) GetEntryById(id int64) (*model.Entry, error) {
	entry, err := x.queries.GetEntryById(context.Background(), id)

	if err == sql.ErrNoRows {
		return nil, goerr.Wrap(types.ErrRecordNotFound).With("entry_id", id)
	}

	if err != nil {
		return nil, goerr.Wrap(err, "Failed query to fetch entry").With("entry_id", id)
	}

	return &model.Entry{
		ID:       entry.ID,
		Day:      entry.Day,
		Title:    entry.Title,
		Comment:  entry.Comment,
		URL:      entry.Url,
		ImageURL: entry.ImageUrl,
		Owner: &model.User{
			ID:      entry.UserID,
			Name:    entry.UserName,
			IconURL: entry.UserIconUrl,
		},
	}, nil
}

type CreateEntryInput struct {
	CalendarID int64
	UserID     int64
	Day        int32
}

func (x *Usecase) CreateEntry(input *CreateEntryInput) (*model.Entry, error) {
	_, err := x.GetCalendarById(input.CalendarID)
	if err != nil {
		return nil, err
	}

	if input.Day < 1 || input.Day > 25 {
		return nil, goerr.Wrap(types.ErrInvalidArgument, "Invalid day").With("day", input.Day)
	}

	lastID, err := x.queries.CreateEntry(context.Background(), adventar_db.CreateEntryParams(*input))
	if err != nil {
		return nil, goerr.Wrap(err, "Failed to insert entry")
	}

	return x.GetEntryById(lastID)
}

type DeleteEntryInput struct {
	EntryID int64
	UserID  int64
}

func (x *Usecase) DeleteEntry(input *DeleteEntryInput) error {
	deletable, err := x.entryDeletable(input.EntryID, input.UserID)

	if err != nil {
		return err
	}

	if deletable == false {
		return goerr.Wrap(types.ErrPermissionDenied, "Invalid request")
	}

	err = x.queries.DeleteEntry(context.Background(), input.EntryID)

	if err != nil {
		return goerr.Wrap(err, "Failed to delete entry")
	}

	return nil
}

// ユーザーがエントリを削除可能かどうかを判定する
// エントリの所有者かカレンダーの所有者であれば削除可能
func (x *Usecase) entryDeletable(entryID int64, userID int64) (bool, error) {
	result, err := x.queries.GetEntryAndCalendarOwnerByEntryId(context.Background(), entryID)

	if err != nil {
		return false, goerr.Wrap(err, "Failed query")
	}

	if userID == result.EntryOwnerID {
		return true, nil
	}

	if userID == result.CalendarOwnerID {
		return true, nil
	}

	return false, nil
}
