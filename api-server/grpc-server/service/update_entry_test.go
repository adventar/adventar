package service_test

import (
	"context"
	"testing"

	pb "github.com/adventar/adventar/api-server/grpc-server/grpc/adventar/v1"
	"google.golang.org/grpc/metadata"
)

func TestUpdateEntry(t *testing.T) {
	cleanupDatabase()

	u := &user{name: "foo", authUID: "xxx", authProvider: "google"}
	createUser(t, u)

	c := &calendar{title: "a", description: "b", userID: u.id, year: 2019}
	createCalendar(t, c)

	e := &entry{userID: u.id, calendarID: c.id, day: 1}
	createEntry(t, e)

	in := &pb.UpdateEntryRequest{EntryId: e.id, Comment: "comment", Url: "http://example.com"}
	md := make(map[string][]string)
	md["authorization"] = append(md["authorization"], u.authUID)
	ctx := metadata.NewIncomingContext(context.Background(), md)

	entry, err := service.UpdateEntry(ctx, in)
	if err != nil {
		t.Fatal(err)
	}

	if entry.Comment != "comment" {
		t.Errorf("actual: %s, expected: comment", entry.Comment)
	}

	if entry.Url != "http://example.com" {
		t.Errorf("actual: %s, expected: http://example.com", entry.Url)
	}

	if entry.Title != "site title" {
		t.Errorf("actual: %s, expected: site title", entry.Title)
	}

	if entry.ImageUrl != "http://example.com/image" {
		t.Errorf("actual: %s, expected: http://example.com/image", entry.ImageUrl)
	}
}
