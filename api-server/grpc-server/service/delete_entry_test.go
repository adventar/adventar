package service_test

import (
	"context"
	"testing"

	pb "github.com/adventar/adventar/api-server/grpc-server/grpc/adventar/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func TestDeleteEntryWithCalendarOwner(t *testing.T) {
	cleanupDatabase()

	calendarOwner := &user{name: "foo", authUID: "xxx", authProvider: "google"}
	createUser(t, calendarOwner)

	entryOwner := &user{name: "bar", authUID: "yyy", authProvider: "google"}
	createUser(t, entryOwner)

	c := &calendar{title: "a", description: "b", userID: calendarOwner.id, year: 2019}
	createCalendar(t, c)

	e := &entry{userID: entryOwner.id, calendarID: c.id, day: 1}
	createEntry(t, e)

	md := make(map[string][]string)
	md["authorization"] = append(md["authorization"], calendarOwner.authUID)
	ctx := metadata.NewIncomingContext(context.Background(), md)

	_, err := service.DeleteEntry(ctx, &pb.DeleteEntryRequest{EntryId: e.id})
	if err != nil {
		t.Fatal(err)
	}

	var count int
	err = db.Get(&count, "select count(*) from entries")
	if err != nil {
		t.Fatal(err)
	}
	if count != 0 {
		t.Errorf("actual: %d, expected: 0", count)
	}
}

func TestDeleteEntryWithEntryOwner(t *testing.T) {
	cleanupDatabase()

	calendarOwner := &user{name: "foo", authUID: "xxx", authProvider: "google"}
	createUser(t, calendarOwner)

	entryOwner := &user{name: "bar", authUID: "yyy", authProvider: "google"}
	createUser(t, entryOwner)

	c := &calendar{title: "a", description: "b", userID: calendarOwner.id, year: 2019}
	createCalendar(t, c)

	e := &entry{userID: entryOwner.id, calendarID: c.id, day: 1}
	createEntry(t, e)

	md := make(map[string][]string)
	md["authorization"] = append(md["authorization"], entryOwner.authUID)
	ctx := metadata.NewIncomingContext(context.Background(), md)

	_, err := service.DeleteEntry(ctx, &pb.DeleteEntryRequest{EntryId: e.id})
	if err != nil {
		t.Fatal(err)
	}

	var count int
	err = db.Get(&count, "select count(*) from entries")
	if err != nil {
		t.Fatal(err)
	}
	if count != 0 {
		t.Errorf("actual: %d, expected: 0", count)
	}
}

func TestDeleteEntryWithOtherUser(t *testing.T) {
	cleanupDatabase()

	calendarOwner := &user{name: "foo", authUID: "xxx", authProvider: "google"}
	createUser(t, calendarOwner)

	entryOwner := &user{name: "bar", authUID: "yyy", authProvider: "google"}
	createUser(t, entryOwner)

	otherUser := &user{name: "baz", authUID: "zzz", authProvider: "google"}
	createUser(t, otherUser)

	c := &calendar{title: "a", description: "b", userID: calendarOwner.id, year: 2019}
	createCalendar(t, c)

	e := &entry{userID: entryOwner.id, calendarID: c.id, day: 1}
	createEntry(t, e)

	md := make(map[string][]string)
	md["authorization"] = append(md["authorization"], otherUser.authUID)
	ctx := metadata.NewIncomingContext(context.Background(), md)

	_, err := service.DeleteEntry(ctx, &pb.DeleteEntryRequest{EntryId: e.id})
	s, _ := status.FromError(err)
	if s.Code() != codes.PermissionDenied {
		t.Fatal(err)
	}

	var count int
	err = db.Get(&count, "select count(*) from entries")
	if err != nil {
		t.Fatal(err)
	}
	if count != 1 {
		t.Errorf("actual: %d, expected: 1", count)
	}
}
