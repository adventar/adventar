package main;

type calendar struct {
	ID          int64
	UserID      int64
	Title       string
	Description string
	Year        int32
}

type user struct {
	ID      int64
	Name    string
	IconURL string
}