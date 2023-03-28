-- name: GetCalendarById :one
SELECT * FROM calendars WHERE id = ? LIMIT 1;

-- name: ListCalendarsByYear :many
SELECT * FROM calendars WHERE year = ?;
