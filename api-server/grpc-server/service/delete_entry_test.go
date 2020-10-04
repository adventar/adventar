package service_test

import (
	"context"
	"os"
	"testing"

	pb "github.com/adventar/adventar/api-server/grpc-server/grpc/adventar/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func TestDeleteEntry(t *testing.T) {
	cleanupDatabase()

	calendarOwner := &user{name: "foo", authUID: "xxx", authProvider: "google"}
	createUser(t, calendarOwner)

	entryOwner := &user{name: "bar", authUID: "yyy", authProvider: "google"}
	createUser(t, entryOwner)

	c := &calendar{title: "a", description: "b", userID: calendarOwner.id, year: 2019}
	createCalendar(t, c)

	e1 := &entry{userID: entryOwner.id, calendarID: c.id, day: 1}
	createEntry(t, e1)

	e2 := &entry{userID: entryOwner.id, calendarID: c.id, day: 2}
	createEntry(t, e2)

	e3 := &entry{userID: entryOwner.id, calendarID: c.id, day: 3}
	createEntry(t, e3)

	md := make(map[string][]string)
	md["authorization"] = append(md["authorization"], calendarOwner.authUID)
	ctx := metadata.NewIncomingContext(context.Background(), md)

	os.Setenv("CURRENT_DATE", "2019-12-02 23:59:59")
	_, err := service.DeleteEntry(ctx, &pb.DeleteEntryRequest{EntryId: e1.id})
	s, _ := status.FromError(err)
	if s.Code() != codes.PermissionDenied {
		t.Fatal(err)
	}

	os.Setenv("CURRENT_DATE", "2019-12-03 00:00:00")
	_, err = service.DeleteEntry(ctx, &pb.DeleteEntryRequest{EntryId: e1.id})
	if err != nil {
		t.Fatal(err)
	}

	md = make(map[string][]string)
	md["authorization"] = append(md["authorization"], entryOwner.authUID)
	os.Setenv("CURRENT_DATE", "2019-12-01 00:00:00")
	ctx = metadata.NewIncomingContext(context.Background(), md)
	_, err = service.DeleteEntry(ctx, &pb.DeleteEntryRequest{EntryId: e2.id})
	if err != nil {
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
