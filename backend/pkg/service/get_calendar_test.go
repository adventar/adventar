package service_test

import (
	"context"
	"testing"

	adventarv1 "github.com/adventar/adventar/backend/pkg/gen/adventar/v1"
	"github.com/bufbuild/connect-go"
)

func TestGetCalendar(t *testing.T) {
	cleanupDatabase()

	ctx := context.Background()

	u := &user{name: "foo", authUID: "xxx", authProvider: "google"}
	createUser(t, u)

	c := &calendar{title: "a", description: "b", userID: u.id, year: 2019}
	createCalendar(t, c)

	e := &entry{userID: u.id, calendarID: c.id, day: 1}
	createEntry(t, e)

	req := connect.NewRequest(&adventarv1.GetCalendarRequest{
		CalendarId: c.id,
	})

	res, err := service.GetCalendar(ctx, req)
	if err != nil {
		t.Fatalf("%+v", err)
	}

	if res.Msg.Calendar.Id != req.Msg.GetCalendarId() {
		t.Errorf("actual: %d, expected: %d", res.Msg.Calendar.Id, req.Msg.CalendarId)
	}

	if res.Msg.Calendar.EntryCount != 1 {
		t.Errorf("actual: %d, expected: %d", res.Msg.Calendar.EntryCount, 1)
	}

	if res.Msg.Entries[0].Id != e.id {
		t.Errorf("actual: %d, expected: %d", res.Msg.Entries[0].Id, e.id)
	}
}
