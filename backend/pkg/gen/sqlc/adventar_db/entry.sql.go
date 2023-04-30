// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: entry.sql

package adventar_db

import (
	"context"
)

const createEntry = `-- name: CreateEntry :execlastid
INSERT INTO entries
  (calendar_id, user_id, day, comment, url, title, image_url)
VALUES
  (?, ?, ?, '', '', '', '')
`

type CreateEntryParams struct {
	CalendarID int64
	UserID     int64
	Day        int32
}

func (q *Queries) CreateEntry(ctx context.Context, arg CreateEntryParams) (int64, error) {
	result, err := q.db.ExecContext(ctx, createEntry, arg.CalendarID, arg.UserID, arg.Day)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

const deleteEntry = `-- name: DeleteEntry :exec
DELETE FROM
  entries
WHERE
  id = ?
`

func (q *Queries) DeleteEntry(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteEntry, id)
	return err
}

const getEntryAndCalendarOwnerByEntryId = `-- name: GetEntryAndCalendarOwnerByEntryId :one
SELECT
  e.user_id AS entry_owner_id,
  c.user_id AS calendar_owner_id
FROM
  entries AS e
  INNER JOIN calendars AS c ON c.id = e.calendar_id
WHERE
  e.id = ?
`

type GetEntryAndCalendarOwnerByEntryIdRow struct {
	EntryOwnerID    int64
	CalendarOwnerID int64
}

func (q *Queries) GetEntryAndCalendarOwnerByEntryId(ctx context.Context, id int64) (GetEntryAndCalendarOwnerByEntryIdRow, error) {
	row := q.db.QueryRowContext(ctx, getEntryAndCalendarOwnerByEntryId, id)
	var i GetEntryAndCalendarOwnerByEntryIdRow
	err := row.Scan(&i.EntryOwnerID, &i.CalendarOwnerID)
	return i, err
}

const getEntryById = `-- name: GetEntryById :one
SELECT
  entries.id,
  entries.day,
  entries.title,
  entries.comment,
  entries.url,
  entries.image_url,
  users.id as user_id,
  users.name as user_name,
  users.icon_url as user_icon_url
FROM
  entries
INNER JOIN
  users ON entries.user_id = users.id
WHERE
  entries.id = ?
`

type GetEntryByIdRow struct {
	ID          int64
	Day         int32
	Title       string
	Comment     string
	Url         string
	ImageUrl    string
	UserID      int64
	UserName    string
	UserIconUrl string
}

func (q *Queries) GetEntryById(ctx context.Context, id int64) (GetEntryByIdRow, error) {
	row := q.db.QueryRowContext(ctx, getEntryById, id)
	var i GetEntryByIdRow
	err := row.Scan(
		&i.ID,
		&i.Day,
		&i.Title,
		&i.Comment,
		&i.Url,
		&i.ImageUrl,
		&i.UserID,
		&i.UserName,
		&i.UserIconUrl,
	)
	return i, err
}

const listEntriesByCalendarId = `-- name: ListEntriesByCalendarId :many
SELECT
  entries.id,
  entries.day,
  entries.title,
  entries.comment,
  entries.url,
  entries.image_url,
  users.id as user_id,
  users.name as user_name,
  users.icon_url as user_icon_url
FROM
  entries
INNER JOIN
  users ON entries.user_id = users.id
WHERE
  entries.calendar_id = ?
ORDER BY
  entries.day
`

type ListEntriesByCalendarIdRow struct {
	ID          int64
	Day         int32
	Title       string
	Comment     string
	Url         string
	ImageUrl    string
	UserID      int64
	UserName    string
	UserIconUrl string
}

func (q *Queries) ListEntriesByCalendarId(ctx context.Context, calendarID int64) ([]ListEntriesByCalendarIdRow, error) {
	rows, err := q.db.QueryContext(ctx, listEntriesByCalendarId, calendarID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListEntriesByCalendarIdRow
	for rows.Next() {
		var i ListEntriesByCalendarIdRow
		if err := rows.Scan(
			&i.ID,
			&i.Day,
			&i.Title,
			&i.Comment,
			&i.Url,
			&i.ImageUrl,
			&i.UserID,
			&i.UserName,
			&i.UserIconUrl,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
