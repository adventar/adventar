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
