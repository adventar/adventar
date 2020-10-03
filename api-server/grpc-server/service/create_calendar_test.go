package service_test

import (
	"context"
	"os"
	"strings"
	"testing"

	pb "github.com/adventar/adventar/api-server/grpc-server/grpc/adventar/v1"
	"google.golang.org/grpc/metadata"
)

func TestCreateCalendar(t *testing.T) {
	os.Setenv("CURRENT_DATE", "2019-11-01 00:00:00")
	cleanupDatabase()

	u := &user{name: "foo", authUID: "xxx", authProvider: "google"}
	createUser(t, u)

	in := &pb.CreateCalendarRequest{Title: "foo", Description: "bar"}
	md := make(map[string][]string)
	md["authorization"] = append(md["authorization"], u.authUID)
	ctx := metadata.NewIncomingContext(context.Background(), md)

	calendar, err := service.CreateCalendar(ctx, in)
	if err != nil {
		t.Fatal(err)
	}

	var count int
	err = db.QueryRow("select count(*) from calendars").Scan(&count)
	if err != nil {
		t.Fatal(err)
	}
	if count != 1 {
		t.Errorf("actual: %d, expected: 1", count)
	}

	var title string
	err = db.QueryRow("select title from calendars where id = ?", calendar.Id).Scan(&title)
	if err != nil {
		t.Fatal(err)
	}
	if title != "foo" {
		t.Errorf("actual: %s, expected: foo", title)
	}
}

func TestCalendarCreatable(t *testing.T) {
	in := &pb.CreateCalendarRequest{Title: "foo", Description: "bar"}
	ctx := context.Background()

	os.Setenv("CURRENT_DATE", "2019-10-31 23:59:59")
	_, err := service.CreateCalendar(ctx, in)
	if err == nil || !strings.Contains(err.Error(), "Calendars can not create now") {
		t.Errorf("Unexpected error: %s", err)
	}

	os.Setenv("CURRENT_DATE", "2019-01-01 00:00:00")
	_, err = service.CreateCalendar(ctx, in)
	if err == nil || !strings.Contains(err.Error(), "Calendars can not create now") {
		t.Errorf("Unexpected error: %s", err)
	}
}
