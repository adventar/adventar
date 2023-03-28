// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: user.sql

package adventar_db

import (
	"context"
)

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
