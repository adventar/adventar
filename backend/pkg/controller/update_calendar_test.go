package controller_test

import (
	"context"
	"testing"

	"github.com/adventar/adventar/backend/pkg/controller"
	adventarv1 "github.com/adventar/adventar/backend/pkg/gen/proto/adventar/v1"
	"github.com/bufbuild/connect-go"
)

func TestUpdateCalendar(t *testing.T) {
	cleanupDatabase()

	u := &user{name: "foo", authUID: "xxx", authProvider: "google"}
	createUser(t, u)

	c := &calendar{title: "a", description: "b", userID: u.id, year: 2019}
	createCalendar(t, c)

	req := connect.NewRequest(&adventarv1.UpdateCalendarRequest{
		CalendarId:  c.id,
		Title:       "foo",
		Description: "bar",
	})
	req.Header().Set("authorization", u.authUID)
	ctx := controller.SetRequestMetadata(context.Background(), req)
	_, err := service.UpdateCalendar(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	var r struct {
		Title       string
		Description string
	}
	err = db.Get(&r, "select title, description from calendars where id = ?", c.id)
	if err != nil {
		t.Fatal(err)
	}
	if r.Title != "foo" {
		t.Errorf("actual: %s, expected: foo", r.Title)
	}
	if r.Description != "bar" {
		t.Errorf("actual: %s, expected: bar", r.Description)
	}
}
