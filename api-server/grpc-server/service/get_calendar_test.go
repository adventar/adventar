package service_test

import (
	"context"
	"testing"

	pb "github.com/adventar/adventar/api-server/grpc-server/grpc/adventar/v1"
)

func TestGetCalendar(t *testing.T) {
	cleanupDatabase()

	in := new(pb.GetCalendarRequest)
	ctx := context.Background()

	u := &user{name: "foo", authUID: "xxx", authProvider: "google"}
	createUser(t, u)

	c := &calendar{title: "a", description: "b", userID: u.id, year: 2019}
	createCalendar(t, c)

	e := &entry{userID: u.id, calendarID: c.id, day: 1}
	createEntry(t, e)

	in.CalendarId = c.id

	res, err := service.GetCalendar(ctx, in)
	if err != nil {
		t.Fatalf("%+v", err)
	}

	if res.Calendar.Id != in.CalendarId {
		t.Errorf("actual: %d, expected: %d", res.Calendar.Id, in.CalendarId)
	}

	if res.Calendar.EntryCount != 1 {
		t.Errorf("actual: %d, expected: %d", res.Calendar.EntryCount, 1)
	}

	if res.Entries[0].Id != e.id {
		t.Errorf("actual: %d, expected: %d", res.Entries[0].Id, e.id)
	}
}
