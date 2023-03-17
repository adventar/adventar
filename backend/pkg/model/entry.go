package model

import "time"

// Entry model
type Entry struct {
	ID         int64     `db:"id"`
	UserID     int64     `db:"user_id"`
	CalendarID int64     `db:"calendar_id"`
	Day        int32     `db:"day"`
	Comment    string    `db:"comment"`
	URL        string    `db:"url"`
	Title      string    `db:"title"`
	ImageURL   string    `db:"image_url"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}
