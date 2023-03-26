package service

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	adventarv1 "github.com/adventar/adventar/backend/pkg/gen/adventar/v1"
	"github.com/adventar/adventar/backend/pkg/model"
	"github.com/bufbuild/connect-go"
	"golang.org/x/xerrors"
)

// ListCalendars lists calendars.
func (s *Service) ListCalendars(
	ctx context.Context,
	req *connect.Request[adventarv1.ListCalendarsRequest],
) (*connect.Response[adventarv1.ListCalendarsResponse], error) {
	relation := sq.
		Select(makeSelectValue(map[string][]string{
			"calendars": {"id", "title", "description", "year"},
			"users":     {"id", "name", "icon_url"},
		})...).
		From("calendars").
		Where(sq.Eq{"calendars.year": req.Msg.GetYear()}).
		Where(sq.Eq{"listable": true}).
		Join("users on users.id = calendars.user_id").
		OrderBy("calendars.id desc")

	if req.Msg.GetUserId() != 0 {
		relation = relation.Where(sq.Eq{"calendars.user_id": req.Msg.GetUserId()})
	}
	if req.Msg.GetQuery() != "" {
		v := fmt.Sprint("%", req.Msg.GetQuery(), "%")
		relation = relation.Where("(calendars.title like ? or calendars.description like ?)", v, v)
	}
	if req.Msg.GetPageSize() != 0 {
		relation = relation.Limit(uint64(req.Msg.GetPageSize()))
	}

	query, args, err := relation.ToSql()
	if err != nil {
		return nil, xerrors.Errorf("Failed query to create sql: %w", err)
	}

	rows := []struct {
		Calendar model.Calendar `db:"calendars"`
		User     model.User     `db:"users"`
	}{}

	err = s.db.Select(&rows, query, args...)
	if err != nil {
		return nil, xerrors.Errorf("Failed query to fetch entries: %w", err)
	}

	var calendars []*adventarv1.Calendar
	for _, r := range rows {
		calendars = append(calendars, &adventarv1.Calendar{
			Id:          r.Calendar.ID,
			Title:       r.Calendar.Title,
			Description: r.Calendar.Description,
			Year:        r.Calendar.Year,
			Owner: &adventarv1.User{
				Id:      r.User.ID,
				Name:    r.User.Name,
				IconUrl: r.User.IconURL,
			},
		})
	}

	if len(calendars) != 0 {
		err := s.bindEntryCount(calendars)
		if err != nil {
			return nil, xerrors.Errorf("Failed to bind entry count: %w", err)
		}
	}

	return connect.NewResponse(&adventarv1.ListCalendarsResponse{Calendars: calendars}), nil
}

func (s *Service) bindEntryCount(calendars []*adventarv1.Calendar) error {
	ids := []int64{}
	for _, c := range calendars {
		ids = append(ids, c.Id)
	}

	query, args, err := sq.
		Select("calendar_id as cid, count(*) as count").
		From("entries").
		Where(sq.Eq{"calendar_id": ids}).
		GroupBy("cid").
		ToSql()
	if err != nil {
		return xerrors.Errorf("Failed query to create sql: %w", err)
	}

	rows := []struct {
		Cid   int64
		Count int32
	}{}

	err = s.db.Select(&rows, query, args...)
	if err != nil {
		return xerrors.Errorf("Failed query to fetch entry counts: %w", err)
	}

	entryCounts := map[int64]int32{}
	for _, r := range rows {
		entryCounts[r.Cid] = r.Count
	}

	for _, c := range calendars {
		c.EntryCount = entryCounts[c.Id]
	}

	return nil
}
