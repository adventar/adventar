package usecase

import (
	"context"
	"database/sql"

	"github.com/adventar/adventar/backend/pkg/domain/model"
	"github.com/adventar/adventar/backend/pkg/domain/types"
	"github.com/adventar/adventar/backend/pkg/gen/sqlc/adventar_db"
	"github.com/m-mizutani/goerr"
	"github.com/m-mizutani/gots/slice"
)

func (x *Usecase) GetCalendarById(id int64) (*model.Calendar, error) {
	calendar, err := x.queries.GetCalendarWithUserById(context.Background(), id)

	if err == sql.ErrNoRows {
		return nil, types.ErrRecordNotFound.With("calendar_id", id)
	}

	if err != nil {
		return nil, goerr.Wrap(err, "Failed query to fetch calendar").With("calendar_id", id)
	}

	entries, err := x.queries.ListEntriesByCalendarId(context.Background(), id)

	if err != nil {
		return nil, goerr.Wrap(err, "Failed to find entries").With("calendar_id", id)
	}

	return &model.Calendar{
		ID:          calendar.ID,
		Title:       calendar.Title,
		Description: calendar.Description,
		Year:        calendar.Year,
		EntryCount:  int32(len(entries)),
		Owner: &model.User{
			ID:      calendar.UserID,
			Name:    calendar.UserName,
			IconURL: calendar.UserIconUrl,
		},
		Entries: slice.Map(entries, func(entry adventar_db.ListEntriesByCalendarIdRow) *model.Entry {
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
			}
		}),
	}, nil
}
