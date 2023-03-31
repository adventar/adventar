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

-- name: CreateCalendar :execlastid
INSERT INTO calendars
  (title, description, year, user_id)
VALUES
  (?, ?, ?, ?);
