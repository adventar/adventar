package usecase

import (
	"context"
	"database/sql"
	"net/url"
	"strings"

	"github.com/adventar/adventar/backend/pkg/domain/model"
	"github.com/adventar/adventar/backend/pkg/domain/types"
	"github.com/adventar/adventar/backend/pkg/gen/sqlc/adventar_db"
	"github.com/adventar/adventar/backend/pkg/util"
	"github.com/m-mizutani/goerr"
	"golang.org/x/exp/slog"
)

func (x *Usecase) GetEntryById(id int64) (*model.Entry, error) {
	entry, err := x.queries.GetEntryById(context.Background(), id)

	if err == sql.ErrNoRows {
		return nil, goerr.Wrap(types.ErrRecordNotFound, "Entry not found").With("entry_id", id)
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

func (x *Usecase) ListUserEntries(userID int64) ([]*model.Entry, error) {
	entries, err := x.queries.ListUserEntries(context.Background(), userID)

	if err != nil {
		return nil, goerr.Wrap(err, "Failed query to fetch entries").With("user_id", userID)
	}

	var results []*model.Entry
	for _, entry := range entries {
		results = append(results, &model.Entry{
			ID:       entry.ID,
			Day:      entry.Day,
			Title:    entry.Title,
			Comment:  entry.Comment,
			URL:      entry.Url,
			ImageURL: entry.ImageUrl,
			Calendar: &model.Calendar{
				ID:          entry.CalendarID,
				Title:       entry.CalendarTitle,
				Description: entry.CalendarDescription,
				Year:        entry.CalendarYear,
			},
			Owner: &model.User{
				ID:      entry.UserID,
				Name:    entry.UserName,
				IconURL: entry.UserIconUrl,
			},
		})
	}

	return results, nil
}

func (x *Usecase) ListUserEntriesByYear(userID int64, year int32) ([]*model.Entry, error) {
	params := adventar_db.ListUserEntriesByYearParams{
		UserID: userID,
		Year:   year,
	}
	entries, err := x.queries.ListUserEntriesByYear(context.Background(), params)

	if err != nil {
		return nil, goerr.Wrap(err, "Failed query to fetch entries").With("user_id", userID).With("year", year)
	}

	var results []*model.Entry
	for _, entry := range entries {
		results = append(results, &model.Entry{
			ID:       entry.ID,
			Day:      entry.Day,
			Title:    entry.Title,
			Comment:  entry.Comment,
			URL:      entry.Url,
			ImageURL: entry.ImageUrl,
			Calendar: &model.Calendar{
				ID:          entry.CalendarID,
				Title:       entry.CalendarTitle,
				Description: entry.CalendarDescription,
				Year:        entry.CalendarYear,
			},
			Owner: &model.User{
				ID:      entry.UserID,
				Name:    entry.UserName,
				IconURL: entry.UserIconUrl,
			},
		})
	}

	return results, nil
}

type CreateEntryInput struct {
	CalendarID int64
	UserID     int64
	Day        int32
}

func (x *Usecase) CreateEntry(input *CreateEntryInput) (*model.Entry, error) {
	_, err := x.GetCalendarById(input.CalendarID)
	if err != nil {
		return nil, goerr.Wrap(err, "Failed to get calendar")
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

type UpdateEntryInput struct {
	EntryID int64
	UserID  int64
	URL     string
	Comment string
}

func (x *Usecase) UpdateEntry(input *UpdateEntryInput) (*model.Entry, error) {
	url := strings.TrimSpace(input.URL)
	if url != "" && !isValidURL(url) {
		return nil, goerr.Wrap(types.ErrInvalidArgument, "Invalid URL").With("url", url)
	}

	params := adventar_db.UpdateEntryCommentAndUrlParams{
		ID:      input.EntryID,
		UserID:  input.UserID,
		Url:     input.URL,
		Comment: input.Comment,
	}
	err := x.queries.UpdateEntryCommentAndUrl(context.Background(), params)

	if err != nil {
		return nil, goerr.Wrap(err, "Failed to update entry").With("entry_id", input.EntryID)
	}

	if input.URL != "" {
		m, err := x.metaFetcher.Fetch(input.URL)
		var title string
		var imageURL string
		if err != nil {
			util.Logger.Warn("Failed to fetch site meta info",
				slog.String("url", input.URL),
				slog.Any("err", err.Error()),
			)
			title = ""
			imageURL = ""
		} else {
			title = m.Title
			imageURL = m.ImageURL
		}

		params := adventar_db.UpdateEntryTitleAndImageUrlParams{
			ID:       input.EntryID,
			UserID:   input.UserID,
			Title:    title,
			ImageUrl: imageURL,
		}
		err = x.queries.UpdateEntryTitleAndImageUrl(context.Background(), params)

		if err != nil {
			return nil, goerr.Wrap(err, "Failed query to update entry").With("entry_id", input.EntryID)
		}
	}

	return x.GetEntryById(input.EntryID)
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

func isValidURL(s string) bool {
	u, err := url.Parse(s)
	if err != nil {
		return false
	}

	return u.Scheme == "http" || u.Scheme == "https"
}
