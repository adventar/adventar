package service_test

import (
	"context"
	"os"
	"strings"
	"testing"

	adventarv1 "github.com/adventar/adventar/backend/pkg/gen/proto/adventar/v1"
	s "github.com/adventar/adventar/backend/pkg/service"
	"github.com/bufbuild/connect-go"
)

func TestCreateCalendar(t *testing.T) {
	os.Setenv("CURRENT_DATE", "2019-11-01 00:00:00")
	cleanupDatabase()

	u := &user{name: "foo", authUID: "xxx", authProvider: "google"}
	createUser(t, u)

	req := connect.NewRequest(&adventarv1.CreateCalendarRequest{Title: "foo", Description: "bar"})
	req.Header().Set("authorization", u.authUID)
	ctx := s.SetRequestMetadata(context.Background(), req)

	res, err := service.CreateCalendar(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	var count int
	err = db.Get(&count, "select count(*) from calendars")
	if err != nil {
		t.Fatal(err)
	}
	if count != 1 {
		t.Errorf("actual: %d, expected: 1", count)
	}

	var title string
	err = db.Get(&title, "select title from calendars where id = ?", res.Msg.Id)
	if err != nil {
		t.Fatal(err)
	}
	if title != "foo" {
		t.Errorf("actual: %s, expected: foo", title)
	}
}

func TestCalendarCreatable(t *testing.T) {
	cleanupDatabase()

	u := &user{name: "foo", authUID: "xxx", authProvider: "google"}
	createUser(t, u)
	req := connect.NewRequest(&adventarv1.CreateCalendarRequest{Title: "foo", Description: "bar"})
	req.Header().Set("authorization", u.authUID)
	ctx := s.SetRequestMetadata(context.Background(), req)

	os.Setenv("CURRENT_DATE", "2019-10-31 23:59:59")
	_, err := service.CreateCalendar(ctx, req)
	if err == nil || !strings.Contains(err.Error(), "Calendars can not create now") {
		t.Errorf("Unexpected error: %s", err)
	}

	os.Setenv("CURRENT_DATE", "2019-01-01 00:00:00")
	_, err = service.CreateCalendar(ctx, req)
	if err == nil || !strings.Contains(err.Error(), "Calendars can not create now") {
		t.Errorf("Unexpected error: %s", err)
	}
}
