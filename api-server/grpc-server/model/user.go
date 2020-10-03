package model

// User model
type User struct {
	ID      int64  `db:"id"`
	Name    string `db:"name"`
	IconURL string `db:"icon_url"`
}
