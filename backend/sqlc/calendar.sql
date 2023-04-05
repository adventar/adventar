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

-- name: ListCalendarsByYear :many
SELECT * FROM calendars WHERE year = ?;

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

-- name: CreateCalendar :execlastid
INSERT INTO calendars
  (title, description, year, user_id)
VALUES
  (?, ?, ?, ?);

-- name: DeleteCalendar :exec
DELETE FROM
  calendars
WHERE
  id = ?
  AND user_id = ?
