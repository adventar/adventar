package usecase

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/adventar/adventar/backend/pkg/domain/model"
	"github.com/adventar/adventar/backend/pkg/domain/types"
	"github.com/adventar/adventar/backend/pkg/gen/sqlc/adventar_db"
	"github.com/adventar/adventar/backend/pkg/util"
	"github.com/m-mizutani/goerr"
	"github.com/m-mizutani/gots/slice"
)

type calendarRow struct {
	ID          int64
	Title       string
	Description string
	Year        int32
	UserID      int64
	UserName    string
	UserIconUrl string
}

type calendarRowToModelInput struct {
	entryCount int32
	entries    []adventar_db.ListEntriesByCalendarIdRow
}

func (r *calendarRow) toModel(input *calendarRowToModelInput) *model.Calendar {
	entries := input.entries
	if entries == nil {
		entries = []adventar_db.ListEntriesByCalendarIdRow{}
	}

	return &model.Calendar{
		ID:          r.ID,
		Title:       r.Title,
		Description: r.Description,
		Year:        r.Year,
		EntryCount:  input.entryCount,
		Owner: &model.User{
			ID:      r.UserID,
			Name:    r.UserName,
			IconURL: r.UserIconUrl,
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
	}
}

func (x *Usecase) GetCalendarById(id int64) (*model.Calendar, error) {
	calendar, err := x.queries.GetCalendarWithUserById(context.Background(), id)

	if err == sql.ErrNoRows {
		return nil, goerr.Wrap(types.ErrRecordNotFound, "Calendar not found").With("calendar_id", id)
	}

	if err != nil {
		return nil, goerr.Wrap(err, "Failed query to fetch calendar").With("calendar_id", id)
	}

	entries, err := x.queries.ListEntriesByCalendarId(context.Background(), id)

	if err != nil {
		return nil, goerr.Wrap(err, "Failed to find entries").With("calendar_id", id)
	}

	row := calendarRow(calendar)

	return row.toModel(&calendarRowToModelInput{
		entryCount: int32(len(entries)),
		entries:    entries,
	}), nil
}

func (x *Usecase) ListCalendars(year int32, limit int32) ([]*model.Calendar, error) {
	calendars, err := x.queries.ListCalendars(context.Background(), adventar_db.ListCalendarsParams{
		Year:  year,
		Limit: limit,
	})

	if err != nil {
		return nil, goerr.Wrap(err, "Failed query to fetch calendars")
	}

	rows := slice.Map(calendars, func(calendar adventar_db.ListCalendarsRow) calendarRow {
		return calendarRow(calendar)
	})

	return x.calendarRowsToModels(rows)
}

func (x *Usecase) ListAllCalendars(year int32) ([]*model.Calendar, error) {
	calendars, err := x.queries.ListAllCalendars(context.Background(), year)

	if err != nil {
		return nil, goerr.Wrap(err, "Failed query to fetch calendars")
	}

	rows := slice.Map(calendars, func(calendar adventar_db.ListAllCalendarsRow) calendarRow {
		return calendarRow(calendar)
	})

	return x.calendarRowsToModels(rows)
}

func (x *Usecase) SearchCalendars(year int32, query string) ([]*model.Calendar, error) {
	calendars, err := x.queries.SearchCalendars(context.Background(), adventar_db.SearchCalendarsParams{
		Year:    year,
		Keyword: fmt.Sprint("%", query, "%"),
	})

	if err != nil {
		return nil, goerr.Wrap(err, "Failed query to fetch calendars")
	}

	rows := slice.Map(calendars, func(calendar adventar_db.SearchCalendarsRow) calendarRow {
		return calendarRow(calendar)
	})

	return x.calendarRowsToModels(rows)
}

func (x *Usecase) ListCalendarsByUserId(year int32, userID int64) ([]*model.Calendar, error) {
	calendars, err := x.queries.ListCalendarsByUserId(context.Background(), adventar_db.ListCalendarsByUserIdParams{
		Year: year,
		ID:   userID,
	})

	if err != nil {
		return nil, goerr.Wrap(err, "Failed query to fetch calendars")
	}

	rows := slice.Map(calendars, func(calendar adventar_db.ListCalendarsByUserIdRow) calendarRow {
		return calendarRow(calendar)
	})

	return x.calendarRowsToModels(rows)
}

func (x *Usecase) ListCalendarStats() ([]*model.CalendarStat, error) {
	stats, err := x.queries.ListCalendarStats(context.Background())

	if err != nil {
		return nil, goerr.Wrap(err, "Failed query to fetch calendars stats")
	}

	return slice.Map(stats, func(stat adventar_db.ListCalendarStatsRow) *model.CalendarStat {
		return &model.CalendarStat{
			Year:           stat.Year,
			CalendarsCount: int32(stat.CalendarsCount),
			EntriesCount:   int32(stat.EntriesCount),
		}
	}), nil
}

type CreateCalendarInput struct {
	Title       string
	Description string
	UserID      int64
}

func (x *Usecase) CreateCalendar(input *CreateCalendarInput) (*model.Calendar, error) {
	now, err := util.CurrentDate()
	if err != nil {
		return nil, goerr.Wrap(err, "Failed to get current date")
	}

	if now.Month < 11 {
		return nil, goerr.Wrap(types.ErrFailedPrecondition, "Calendars can not create now").With("current_date", now)
	}

	if input.Title == "" {
		return nil, goerr.Wrap(types.ErrInvalidArgument, "Title is required")
	}

	lastID, err := x.queries.CreateCalendar(context.Background(), adventar_db.CreateCalendarParams{
		Title:       input.Title,
		Description: input.Description,
		UserID:      input.UserID,
		Year:        int32(now.Year),
	})

	if err != nil {
		return nil, goerr.Wrap(err, "Failed to insert calendar").With("input", input)
	}

	return x.GetCalendarById(lastID)
}

type UpdateCalendarInput struct {
	CalendarID  int64
	Title       string
	Description string
	UserID      int64
}

func (x *Usecase) UpdateCalendar(input *UpdateCalendarInput) (*model.Calendar, error) {
	if input.Title == "" {
		return nil, goerr.Wrap(types.ErrInvalidArgument, "Title is required")
	}

	err := x.queries.UpdateCalendar(context.Background(), adventar_db.UpdateCalendarParams{
		ID:          input.CalendarID,
		Title:       input.Title,
		Description: input.Description,
		UserID:      input.UserID,
	})

	if err != nil {
		return nil, goerr.Wrap(err, "Failed query to update calendar")
	}

	return x.GetCalendarById(input.CalendarID)
}

type DeleteCalendarInput struct {
	CalendarID int64
	UserID     int64
}

func (x *Usecase) DeleteCalendar(input *DeleteCalendarInput) error {
	err := x.queries.DeleteCalendar(context.Background(), adventar_db.DeleteCalendarParams{
		ID:     input.CalendarID,
		UserID: input.UserID,
	})

	if err != nil {
		return goerr.Wrap(err, "Failed query to delete calendar").With("input", input)
	}

	return nil
}

func (x *Usecase) calendarRowsToModels(rows []calendarRow) ([]*model.Calendar, error) {
	entryCountByCalendarID, err := x.getEntryCounts(rows)

	if err != nil {
		return nil, goerr.Wrap(err, "Failed to get entry counts")
	}

	return slice.Map(rows, func(row calendarRow) *model.Calendar {
		return row.toModel(&calendarRowToModelInput{
			entryCount: entryCountByCalendarID[row.ID],
		})
	}), nil
}

// getEntryCounts returns entry counts by calendar ID.
func (x *Usecase) getEntryCounts(calendars []calendarRow) (map[int64]int32, error) {
	ids := slice.Map(calendars, func(c calendarRow) int64 {
		return c.ID
	})

	result, err := x.queries.GetEntryCountByCalendarId(context.Background(), ids)
	if err != nil {
		return nil, goerr.Wrap(err, "Failed to GetEntryCountByCalendarId query")
	}

	counts := map[int64]int32{}
	for _, r := range result {
		counts[r.CalendarID] = int32(r.Count)
	}

	return counts, nil
}
