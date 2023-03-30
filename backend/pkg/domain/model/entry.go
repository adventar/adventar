package model

import (
	adventarv1 "github.com/adventar/adventar/backend/pkg/gen/proto/adventar/v1"
	"github.com/adventar/adventar/backend/pkg/util"
)

type Entry struct {
	ID         int64
	UserID     int64
	CalendarID int64
	Day        int32
	Comment    string
	URL        string
	Title      string
	ImageURL   string
	Owner      *User
}

func (x *Entry) ResizableImageURL() string {
	return util.ResizableImageURL(x.ImageURL)
}

func (x *Entry) ToProto() *adventarv1.Entry {
	return &adventarv1.Entry{
		Id:       x.ID,
		Day:      x.Day,
		Title:    x.Title,
		Comment:  x.Comment,
		Url:      x.URL,
		ImageUrl: x.ResizableImageURL(),
		Owner:    x.Owner.ToProto(),
	}
}
