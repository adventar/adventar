-- name: GetUserById :one
SELECT * FROM users WHERE id = ? LIMIT 1;
