-- name: GetCalendarWithUserById :one
SELECT
  calendars.id,
  calendars.title,
  calendars.description,
  calendars.year,
  users.id as user_id,
  users.name as user_name,
  users.icon_url as user_icon_url
FROM
  calendars
INNER JOIN
  users ON calendars.user_id = users.id
WHERE
  calendars.id = ?
LIMIT 1;

-- name: ListCalendars :many
SELECT
  calendars.id,
  calendars.title,
  calendars.description,
  calendars.year,
  users.id as user_id,
  users.name as user_name,
  users.icon_url as user_icon_url
FROM
  calendars
INNER JOIN
  users ON calendars.user_id = users.id
WHERE
  year = ?
  AND listable = true
ORDER BY
  calendars.id DESC
LIMIT ?;

-- name: ListAllCalendars :many
SELECT
  calendars.id,
  calendars.title,
  calendars.description,
  calendars.year,
  users.id as user_id,
  users.name as user_name,
  users.icon_url as user_icon_url
FROM
  calendars
INNER JOIN
  users ON calendars.user_id = users.id
WHERE
  year = ?
  AND listable = true
ORDER BY
  calendars.id DESC;

-- name: SearchCalendars :many
SELECT
  calendars.id,
  calendars.title,
  calendars.description,
  calendars.year,
  users.id as user_id,
  users.name as user_name,
  users.icon_url as user_icon_url
FROM
  calendars
INNER JOIN
  users ON calendars.user_id = users.id
WHERE
  year = ?
  AND listable = true
  AND (calendars.title LIKE sqlc.arg(keyword) OR calendars.description LIKE sqlc.arg(keyword))
ORDER BY
  calendars.id DESC;

-- name: ListCalendarsByUserId :many
SELECT
  calendars.id,
  calendars.title,
  calendars.description,
  calendars.year,
  users.id as user_id,
  users.name as user_name,
  users.icon_url as user_icon_url
FROM
  calendars
INNER JOIN
  users ON calendars.user_id = users.id
WHERE
  year = ?
  AND users.id = ?
ORDER BY
  calendars.id DESC
;

-- name: ListCalendarStats :many
SELECT
  year
  , count(distinct calendars.id) AS calendars_count
  , count(entries.id) AS entries_count
FROM
  calendars
  LEFT JOIN entries ON entries.calendar_id = calendars.id
GROUP BY
  year
ORDER BY
  year DESC;

-- name: GetEntryCountByCalendarId :many
SELECT
  calendar_id,
  count(*) as count
FROM
  entries
WHERE
  calendar_id IN (sqlc.slice('ids'))
GROUP BY
  calendar_id;

-- name: CreateCalendar :execlastid
INSERT INTO calendars
  (title, description, year, user_id)
VALUES
  (?, ?, ?, ?);

-- name: UpdateCalendar :exec
UPDATE
  calendars
SET
  title = ?,
  description = ?
WHERE
  id = ?
  AND user_id = ?;

-- name: DeleteCalendar :exec
DELETE FROM
  calendars
WHERE
  id = ?
  AND user_id = ?
