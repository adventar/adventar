// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: user.sql

package adventar_db

import (
	"context"
)

const getUserByAuthInfo = `-- name: GetUserByAuthInfo :one
SELECT id, name, auth_uid, auth_provider, icon_url, created_at, updated_at FROM users WHERE auth_provider = ? and auth_uid = ? LIMIT 1
`

type GetUserByAuthInfoParams struct {
	AuthProvider string
	AuthUid      string
}

func (q *Queries) GetUserByAuthInfo(ctx context.Context, arg GetUserByAuthInfoParams) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByAuthInfo, arg.AuthProvider, arg.AuthUid)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.AuthUid,
		&i.AuthProvider,
		&i.IconUrl,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserById = `-- name: GetUserById :one
SELECT id, name, auth_uid, auth_provider, icon_url, created_at, updated_at FROM users WHERE id = ? LIMIT 1
`

func (q *Queries) GetUserById(ctx context.Context, id int64) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserById, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.AuthUid,
		&i.AuthProvider,
		&i.IconUrl,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :exec
UPDATE users SET name = ?  where id = ?
`

type UpdateUserParams struct {
	Name string
	ID   int64
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.db.ExecContext(ctx, updateUser, arg.Name, arg.ID)
	return err
}
