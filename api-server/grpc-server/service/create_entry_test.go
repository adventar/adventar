package service_test

import (
	"context"
	"testing"

	pb "github.com/adventar/adventar/api-server/grpc-server/grpc/adventar/v1"
	"google.golang.org/grpc/metadata"
)

func TestCreateEntry(t *testing.T) {
	cleanupDatabase()

	u := &user{name: "foo", authUID: "xxx", authProvider: "google"}
	createUser(t, u)

	c := &calendar{title: "a", description: "b", userID: u.id, year: 2019}
	createCalendar(t, c)

	in := &pb.CreateEntryRequest{CalendarId: c.id, Day: 1}
	md := make(map[string][]string)
	md["authorization"] = append(md["authorization"], u.authUID)
	ctx := metadata.NewIncomingContext(context.Background(), md)

	entry, err := service.CreateEntry(ctx, in)
	if err != nil {
		t.Fatal(err)
	}

	var day int32
	var cid int64
	var uid int64
	err = db.QueryRow("select day, calendar_id, user_id from entries where id = ?", entry.Id).Scan(&day, &cid, &uid)
	if err != nil {
		t.Fatal(err)
	}
	if day != 1 {
		t.Errorf("actual: %d, expected: 2019-12-01", day)
	}
	if cid != c.id {
		t.Errorf("actual: %d, expected: %d", c.id, cid)
	}
	if uid != u.id {
		t.Errorf("actual: %d, expected: %d", u.id, uid)
	}
}
