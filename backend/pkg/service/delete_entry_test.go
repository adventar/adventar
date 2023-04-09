package service_test

import (
	"context"
	"errors"
	"testing"

	"github.com/adventar/adventar/backend/pkg/domain/types"
	adventarv1 "github.com/adventar/adventar/backend/pkg/gen/proto/adventar/v1"
	s "github.com/adventar/adventar/backend/pkg/service"
	"github.com/bufbuild/connect-go"
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

	req := connect.NewRequest(&adventarv1.DeleteEntryRequest{EntryId: e.id})
	req.Header().Set("authorization", calendarOwner.authUID)
	ctx := s.SetRequestMetadata(context.Background(), req)

	_, err := service.DeleteEntry(ctx, req)
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

	req := connect.NewRequest(&adventarv1.DeleteEntryRequest{EntryId: e.id})
	req.Header().Set("authorization", entryOwner.authUID)
	ctx := s.SetRequestMetadata(context.Background(), req)

	_, err := service.DeleteEntry(ctx, req)
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

	req := connect.NewRequest(&adventarv1.DeleteEntryRequest{EntryId: e.id})
	req.Header().Set("authorization", otherUser.authUID)
	ctx := s.SetRequestMetadata(context.Background(), req)

	_, err := service.DeleteEntry(ctx, req)
	if errors.Is(err, types.ErrPermissionDenied) == false {
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
