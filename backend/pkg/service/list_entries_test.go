package service_test

import (
	"context"
	"testing"

	adventarv1 "github.com/adventar/adventar/backend/pkg/gen/proto/adventar/v1"
	"github.com/bufbuild/connect-go"
)

func TestListEntries(t *testing.T) {
	cleanupDatabase()

	u1 := &user{name: "foo", authUID: "xxx", authProvider: "google"}
	createUser(t, u1)

	u2 := &user{name: "foo", authUID: "yyy", authProvider: "google"}
	createUser(t, u2)

	c1 := &calendar{title: "a", description: "b", userID: u1.id, year: 2019}
	createCalendar(t, c1)

	c2 := &calendar{title: "a", description: "b", userID: u1.id, year: 2018}
	createCalendar(t, c2)

	e1 := &entry{userID: u1.id, calendarID: c1.id, day: 1}
	createEntry(t, e1)

	e2 := &entry{userID: u1.id, calendarID: c2.id, day: 1}
	createEntry(t, e2)

	e3 := &entry{userID: u2.id, calendarID: c1.id, day: 2}
	createEntry(t, e3)

	req := connect.NewRequest(&adventarv1.ListEntriesRequest{UserId: u1.id, Year: 2019})
	req.Header().Set("authorization", u1.authUID)

	res, err := service.ListEntries(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}

	if len(res.Msg.Entries) != 1 {
		t.Errorf("actual: %d, expected: 1", len(res.Msg.Entries))
	}

	if res.Msg.Entries[0].Owner.Id != u1.id {
		t.Errorf("actual: %d, expected: %d", res.Msg.Entries[0].Owner.Id, u1.id)
	}
}
