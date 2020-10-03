package service_test

import (
	"context"
	"testing"

	pb "github.com/adventar/adventar/api-server/grpc-server/grpc/adventar/v1"
	"google.golang.org/grpc/metadata"
)

func TestUpdateCalendar(t *testing.T) {
	cleanupDatabase()

	u := &user{name: "foo", authUID: "xxx", authProvider: "google"}
	createUser(t, u)

	c := &calendar{title: "a", description: "b", userID: u.id, year: 2019}
	createCalendar(t, c)

	in := &pb.UpdateCalendarRequest{CalendarId: c.id, Title: "foo", Description: "bar"}
	md := make(map[string][]string)
	md["authorization"] = append(md["authorization"], u.authUID)
	ctx := metadata.NewIncomingContext(context.Background(), md)

	_, err := service.UpdateCalendar(ctx, in)
	if err != nil {
		t.Fatal(err)
	}

	var title string
	var description string
	err = db.QueryRow("select title, description from calendars where id = ?", c.id).Scan(&title, &description)
	if err != nil {
		t.Fatal(err)
	}
	if title != "foo" {
		t.Errorf("actual: %s, expected: foo", title)
	}
	if description != "bar" {
		t.Errorf("actual: %s, expected: bar", description)
	}
}
