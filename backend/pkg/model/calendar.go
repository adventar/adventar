package model

import "time"

// Calendar model
type Calendar struct {
	ID          int64     `db:"id"`
	UserID      int64     `db:"user_id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	Year        int32     `db:"year"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}
