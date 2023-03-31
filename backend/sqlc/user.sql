-- name: GetUserById :one
SELECT * FROM users WHERE id = ? LIMIT 1;

-- name: GetUserByAuthInfo :one
SELECT * FROM users WHERE auth_provider = ? and auth_uid = ? LIMIT 1;
