package model

import "time"

// User model
type User struct {
	ID        int64     `db:"id"`
	Name      string    `db:"name"`
	IconURL   string    `db:"icon_url"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
