package controller_test

import (
	"context"
	"testing"

	"github.com/adventar/adventar/backend/pkg/controller"
	adventarv1 "github.com/adventar/adventar/backend/pkg/gen/proto/adventar/v1"
	"github.com/bufbuild/connect-go"
)

func TestUpdateEntry(t *testing.T) {
	cleanupDatabase()

	u := &user{name: "foo", authUID: "xxx", authProvider: "google"}
	createUser(t, u)

	c := &calendar{title: "a", description: "b", userID: u.id, year: 2019}
	createCalendar(t, c)

	e := &entry{userID: u.id, calendarID: c.id, day: 1}
	createEntry(t, e)

	req := connect.NewRequest(&adventarv1.UpdateEntryRequest{
		EntryId: e.id,
		Comment: "comment",
		Url:     "http://example.com",
	})
	req.Header().Set("authorization", u.authUID)

	ctx := controller.SetRequestMetadata(context.Background(), req)
	res, err := service.UpdateEntry(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	if res.Msg.Comment != "comment" {
		t.Errorf("actual: %s, expected: comment", res.Msg.Comment)
	}

	if res.Msg.Url != "http://example.com" {
		t.Errorf("actual: %s, expected: http://example.com", res.Msg.Url)
	}

	if res.Msg.Title != "site title" {
		t.Errorf("actual: %s, expected: site title", res.Msg.Title)
	}

	if res.Msg.ImageUrl != "http://example.com/image" {
		t.Errorf("actual: %s, expected: http://example.com/image", res.Msg.ImageUrl)
	}
}
