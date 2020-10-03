package service_test

import (
	"context"
	"testing"

	pb "github.com/adventar/adventar/api-server/grpc-server/grpc/adventar/v1"
)

func TestListCalendarStats(t *testing.T) {
	cleanupDatabase()

	u1 := &user{name: "foo", authUID: "xxx", authProvider: "google"}
	createUser(t, u1)

	c1 := &calendar{title: "a", description: "b", userID: u1.id, year: 2019}
	createCalendar(t, c1)

	e1 := &entry{userID: u1.id, calendarID: c1.id, day: 1}
	createEntry(t, e1)

	e2 := &entry{userID: u1.id, calendarID: c1.id, day: 2}
	createEntry(t, e2)

	in := &pb.ListCalendarStatsRequest{}
	ctx := context.Background()

	res, err := service.ListCalendarStats(ctx, in)
	if err != nil {
		t.Fatal(err)
	}

	if len(res.CalendarStats) != 1 {
		t.Errorf("actual: %d, expected: 1", len(res.CalendarStats))
	}

	if res.CalendarStats[0].Year != 2019 {
		t.Errorf("actual: %d, expected: %d", res.CalendarStats[0].Year, 2019)
	}

	if res.CalendarStats[0].CalendarsCount != 1 {
		t.Errorf("actual: %d, expected: %d", res.CalendarStats[0].CalendarsCount, 1)
	}

	if res.CalendarStats[0].EntriesCount != 2 {
		t.Errorf("actual: %d, expected: %d", res.CalendarStats[0].EntriesCount, 1)
	}
}
