package model

import adventarv1 "github.com/adventar/adventar/backend/pkg/gen/proto/adventar/v1"

type Calendar struct {
	ID          int64
	Title       string
	Description string
	Year        int32
	Owner       *User
	EntryCount  int32
	Entries     []*Entry
}

func (x *Calendar) ToProto() *adventarv1.Calendar {
	return &adventarv1.Calendar{
		Id:          x.ID,
		Title:       x.Title,
		Description: x.Description,
		Year:        x.Year,
		EntryCount:  x.EntryCount,
		Owner:       x.Owner.ToProto(),
	}
}
