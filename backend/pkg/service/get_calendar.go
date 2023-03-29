package service

import (
	"context"
	"database/sql"
	"errors"

	adventarv1 "github.com/adventar/adventar/backend/pkg/gen/proto/adventar/v1"
	"github.com/bufbuild/connect-go"
	"github.com/m-mizutani/goerr"
)

// GetCalendar returns a calendar.
func (s *Service) GetCalendar(
	ctx context.Context,
	req *connect.Request[adventarv1.GetCalendarRequest],
) (*connect.Response[adventarv1.GetCalendarResponse], error) {
	calendarId := req.Msg.GetCalendarId()
	calendar, err := s.queries().GetCalendarWithUserById(context.Background(), calendarId)

	if err == sql.ErrNoRows {
		return nil, connect.NewError(connect.CodeNotFound, errors.New("Calendar not found"))
	}

	if err != nil {
		return nil, goerr.Wrap(err, "Failed query to fetch calendar").With("calendar_id", calendarId)
	}

	entries, err := s.queries().ListEntriesByCalendarId(context.Background(), calendarId)

	if err != nil {
		return nil, goerr.Wrap(err, "Failed to find entries")
	}

	pbCalendar := &adventarv1.Calendar{
		Id:          calendar.ID,
		Title:       calendar.Title,
		Description: calendar.Description,
		Year:        calendar.Year,
		EntryCount:  int32(len(entries)),
		Owner: &adventarv1.User{
			Id:      calendar.UserID,
			Name:    calendar.UserName,
			IconUrl: calendar.UserIconUrl,
		},
	}

	pbEntries := []*adventarv1.Entry{}
	for _, entry := range entries {
		pbEntries = append(pbEntries, &adventarv1.Entry{
			Id:       entry.ID,
			Day:      entry.Day,
			Title:    entry.Title,
			Comment:  entry.Comment,
			Url:      entry.Url,
			ImageUrl: convertImageURL(entry.ImageUrl),
			Owner: &adventarv1.User{
				Id:      entry.UserID,
				Name:    entry.UserName,
				IconUrl: entry.UserIconUrl,
			},
		})
	}

	return connect.NewResponse(&adventarv1.GetCalendarResponse{
		Calendar: pbCalendar,
		Entries:  pbEntries,
	}), nil
}
