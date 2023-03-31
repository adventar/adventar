// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package adventar_db

import (
	"time"
)

type Calendar struct {
	ID          int64
	UserID      int64
	Title       string
	Description string
	Year        int32
	Listable    bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Entry struct {
	ID         int64
	UserID     int64
	CalendarID int64
	Day        int32
	Comment    string
	Url        string
	Title      string
	ImageUrl   string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type User struct {
	ID           int64
	Name         string
	AuthUid      string
	AuthProvider string
	IconUrl      string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}