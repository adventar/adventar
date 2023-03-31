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