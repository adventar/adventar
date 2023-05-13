-- name: GetEntryById :one
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
  entries.id = ?;

-- name: ListEntriesByCalendarId :many
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
  entries.day;

-- name: CreateEntry :execlastid
INSERT INTO entries
  (calendar_id, user_id, day, comment, url, title, image_url)
VALUES
  (?, ?, ?, '', '', '', '');

-- name: GetEntryAndCalendarOwnerByEntryId :one
SELECT
  e.user_id AS entry_owner_id,
  c.user_id AS calendar_owner_id
FROM
  entries AS e
  INNER JOIN calendars AS c ON c.id = e.calendar_id
WHERE
  e.id = ?;

-- name: ListUserEntriesByYear :many
SELECT
  entries.id as id,
  entries.day as day,
  entries.title as title,
  entries.comment as comment,
  entries.url as url,
  entries.image_url as image_url,
  calendars.id as calendar_id,
  calendars.title as calendar_title,
  calendars.description as calendar_description,
  calendars.year as calendar_year,
  users.id as user_id,
  users.name as user_name,
  users.icon_url as user_icon_url
FROM
  entries
INNER JOIN
  users ON entries.user_id = users.id
INNER JOIN
  calendars ON entries.calendar_id = calendars.id
WHERE
  entries.user_id = ?
  AND calendars.year = ?
ORDER BY
  calendars.year, entries.day, entries.id;

-- name: ListUserEntries :many
SELECT
  entries.id as id,
  entries.day as day,
  entries.title as title,
  entries.comment as comment,
  entries.url as url,
  entries.image_url as image_url,
  calendars.id as calendar_id,
  calendars.title as calendar_title,
  calendars.description as calendar_description,
  calendars.year as calendar_year,
  users.id as user_id,
  users.name as user_name,
  users.icon_url as user_icon_url
FROM
  entries
INNER JOIN
  users ON entries.user_id = users.id
INNER JOIN
  calendars ON entries.calendar_id = calendars.id
WHERE
  entries.user_id = ?
ORDER BY
  calendars.year, entries.day, entries.id;

-- name: DeleteEntry :exec
DELETE FROM
  entries
WHERE
  id = ?;
