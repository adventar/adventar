package controller_test

import (
	"context"
	"testing"

	"github.com/adventar/adventar/backend/pkg/controller"
	adventarv1 "github.com/adventar/adventar/backend/pkg/gen/proto/adventar/v1"
	"github.com/bufbuild/connect-go"
)

func TestDeleteCalendar(t *testing.T) {
	cleanupDatabase()

	u := &user{name: "foo", authUID: "xxx", authProvider: "google"}
	createUser(t, u)

	c := &calendar{title: "a", description: "b", userID: u.id, year: 2019}
	createCalendar(t, c)

	req := connect.NewRequest(&adventarv1.DeleteCalendarRequest{CalendarId: c.id})
	req.Header().Set("authorization", u.authUID)
	ctx := controller.SetRequestMetadata(context.Background(), req)

	_, err := service.DeleteCalendar(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	var count int
	err = db.Get(&count, "select count(*) from calendars")
	if err != nil {
		t.Fatal(err)
	}
	if count != 0 {
		t.Errorf("actual: %d, expected: 0", count)
	}
}
