package service

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	adventarv1 "github.com/adventar/adventar/backend/pkg/gen/adventar/v1"
	"github.com/adventar/adventar/backend/pkg/model"
	"github.com/bufbuild/connect-go"
	"golang.org/x/xerrors"
)

// ListEntries lists entries.
func (s *Service) ListEntries(
	ctx context.Context,
	req *connect.Request[adventarv1.ListEntriesRequest],
) (*connect.Response[adventarv1.ListEntriesResponse], error) {
	relation := sq.
		Select(makeSelectValue(map[string][]string{
			"entries":   {"id", "day", "title", "comment", "url", "image_url"},
			"calendars": {"id", "title", "description", "year"},
			"users":     {"id", "name", "icon_url"},
		})...).
		From("entries").
		Join("users on users.id = entries.user_id").
		Join("calendars on calendars.id = entries.calendar_id").
		Where(sq.Eq{"entries.user_id": req.Msg.GetUserId()}).
		OrderBy("calendars.year, entries.day, entries.id")

	if req.Msg.GetYear() != 0 {
		relation = relation.Where(sq.Eq{"year": req.Msg.GetYear()})
	}

	query, args, err := relation.ToSql()
	if err != nil {
		return nil, xerrors.Errorf("Failed query to create sql: %w", err)
	}

	rows := []struct {
		Entry    model.Entry    `db:"entries"`
		Calendar model.Calendar `db:"calendars"`
		User     model.User     `db:"users"`
	}{}

	err = s.db.Select(&rows, query, args...)
	if err != nil {
		return nil, xerrors.Errorf("Failed query to fetch entries: %w", err)
	}

	entries := []*adventarv1.Entry{}
	for _, r := range rows {
		entries = append(entries, &adventarv1.Entry{
			Id:       r.Entry.ID,
			Day:      r.Entry.Day,
			Title:    r.Entry.Title,
			Comment:  r.Entry.Comment,
			Url:      r.Entry.URL,
			ImageUrl: convertImageURL(r.Entry.ImageURL),
			Calendar: &adventarv1.Calendar{
				Id:          r.Calendar.ID,
				Title:       r.Calendar.Title,
				Description: r.Calendar.Description,
				Year:        r.Calendar.Year,
			},
			Owner: &adventarv1.User{
				Id:      r.User.ID,
				Name:    r.User.Name,
				IconUrl: r.User.IconURL,
			},
		})
	}

	return connect.NewResponse(&adventarv1.ListEntriesResponse{Entries: entries}), nil
}
