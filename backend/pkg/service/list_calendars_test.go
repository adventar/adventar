package service_test

import (
	"context"
	"testing"

	adventarv1 "github.com/adventar/adventar/backend/pkg/gen/proto/adventar/v1"
	"github.com/bufbuild/connect-go"
)

func TestListCalendars(t *testing.T) {
	cleanupDatabase()

	req := connect.NewRequest(&adventarv1.ListCalendarsRequest{Year: 2019})
	ctx := context.Background()

	u := &user{name: "foo", authUID: "xxx", authProvider: "google"}
	createUser(t, u)

	c := &calendar{title: "a", description: "b", userID: u.id, year: 2019}
	createCalendar(t, c)

	e := &entry{userID: u.id, calendarID: c.id, day: 1}
	createEntry(t, e)

	c2 := &calendar{title: "a", description: "b", userID: u.id, year: 2018}
	createCalendar(t, c2)

	res, err := service.ListCalendars(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	if len(res.Msg.Calendars) != 1 {
		t.Errorf("actual: %d, expected: %d", len(res.Msg.Calendars), 1)
	}
	if res.Msg.Calendars[0].GetOwner().GetId() != u.id {
		t.Errorf("actual: %d, expected: %d", res.Msg.Calendars[0].GetEntryCount(), 1)
	}
	if res.Msg.Calendars[0].GetEntryCount() != 1 {
		t.Errorf("actual: %d, expected: %d", res.Msg.Calendars[0].GetEntryCount(), 1)
	}
}

func TestListCalendarsWithQuery(t *testing.T) {
	cleanupDatabase()

	req := connect.NewRequest(&adventarv1.ListCalendarsRequest{Query: "test", Year: 2019})
	ctx := context.Background()

	u := &user{name: "foo", authUID: "xxx", authProvider: "google"}
	createUser(t, u)

	c1 := &calendar{title: "Test title", description: "", userID: u.id, year: 2019}
	createCalendar(t, c1)

	c2 := &calendar{title: "foo", description: "Calendar test", userID: u.id, year: 2019}
	createCalendar(t, c2)

	c3 := &calendar{title: "foo", description: "bar", userID: u.id, year: 2019}
	createCalendar(t, c3)

	res, err := service.ListCalendars(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	if len(res.Msg.Calendars) != 2 {
		t.Errorf("actual: %d, expected: %d", len(res.Msg.Calendars), 2)
	}
}

func TestListCalendarsWithUserId(t *testing.T) {
	cleanupDatabase()

	users := []*user{}
	for _, n := range []string{"u1", "u2", "u3"} {
		u := &user{name: n, authUID: n}
		createUser(t, u)
		c := &calendar{userID: u.id, year: 2019}
		createCalendar(t, c)
		users = append(users, u)
	}

	req := connect.NewRequest(&adventarv1.ListCalendarsRequest{UserId: users[0].id, Year: 2019})
	ctx := context.Background()

	res, err := service.ListCalendars(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	if len(res.Msg.Calendars) != 1 {
		t.Errorf("actual: %d, expected: %d", len(res.Msg.Calendars), 1)
	}
}

func TestListCalendarsWithPageSize(t *testing.T) {
	cleanupDatabase()

	u := &user{name: "foo", authUID: "xxx", authProvider: "google"}
	createUser(t, u)

	for _, title := range []string{"c1", "c2", "c3"} {
		c1 := &calendar{title: title, description: "", userID: u.id, year: 2019}
		createCalendar(t, c1)
	}

	req := connect.NewRequest(&adventarv1.ListCalendarsRequest{Year: 2019, PageSize: 1})
	ctx := context.Background()

	res, err := service.ListCalendars(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	if len(res.Msg.Calendars) != 1 {
		t.Errorf("actual: %d, expected: %d", len(res.Msg.Calendars), 1)
	}
}
