package model

import adventarv1 "github.com/adventar/adventar/backend/pkg/gen/proto/adventar/v1"

type User struct {
	ID      int64
	Name    string
	IconURL string
}

func (x *User) ToProto() *adventarv1.User {
	return &adventarv1.User{
		Id:      x.ID,
		Name:    x.Name,
		IconUrl: x.IconURL,
	}
}
