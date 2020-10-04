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

	var r struct {
		Day int32 `db:"day"`
		Cid int64 `db:"calendar_id"`
		Uid int64 `db:"user_id"`
	}
	err = db.Get(&r, "select day, calendar_id, user_id from entries where id = ?", entry.Id)
	if err != nil {
		t.Fatal(err)
	}
	if r.Day != 1 {
		t.Errorf("actual: %d, expected: 2019-12-01", r.Day)
	}
	if r.Cid != c.id {
		t.Errorf("actual: %d, expected: %d", c.id, r.Cid)
	}
	if r.Uid != u.id {
		t.Errorf("actual: %d, expected: %d", u.id, r.Uid)
	}
}
