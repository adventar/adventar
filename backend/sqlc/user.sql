-- name: GetUserById :one
SELECT * FROM users WHERE id = ?;

-- name: GetUserByAuthInfo :one
SELECT * FROM users WHERE auth_provider = ? and auth_uid = ?;

-- name: UpdateUserName :exec
UPDATE users SET name = ?  WHERE id = ?;

-- name: UpdateUserIconUrl :exec
UPDATE users SET icon_url = ? WHERE id = ?;

-- name: CreateUser :execlastid
INSERT INTO users (name, auth_uid, auth_provider, icon_url)
values (?, ?, ?, ?);
