package service_test

import (
	"context"
	"testing"

	pb "github.com/adventar/adventar/api-server/grpc-server/grpc/adventar/v1"
	"google.golang.org/grpc/metadata"
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

	in := &pb.ListEntriesRequest{UserId: u1.id, Year: 2019}
	md := make(map[string][]string)
	md["authorization"] = append(md["authorization"], u1.authUID)
	ctx := metadata.NewIncomingContext(context.Background(), md)

	res, err := service.ListEntries(ctx, in)
	if err != nil {
		t.Fatal(err)
	}

	if len(res.Entries) != 1 {
		t.Errorf("actual: %d, expected: 1", len(res.Entries))
	}

	if res.Entries[0].Owner.Id != u1.id {
		t.Errorf("actual: %d, expected: %d", res.Entries[0].Owner.Id, u1.id)
	}
}
