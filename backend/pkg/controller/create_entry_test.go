package controller_test

import (
	"context"
	"testing"

	"github.com/adventar/adventar/backend/pkg/controller"
	adventarv1 "github.com/adventar/adventar/backend/pkg/gen/proto/adventar/v1"
	"github.com/bufbuild/connect-go"
)

func TestCreateEntry(t *testing.T) {
	cleanupDatabase()

	u := &user{name: "foo", authUID: "xxx", authProvider: "google"}
	createUser(t, u)

	c := &calendar{title: "a", description: "b", userID: u.id, year: 2019}
	createCalendar(t, c)

	req := connect.NewRequest(&adventarv1.CreateEntryRequest{CalendarId: c.id, Day: 1})
	req.Header().Set("authorization", u.authUID)
	ctx := controller.SetRequestMetadata(context.Background(), req)

	res, err := service.CreateEntry(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	var r struct {
		Day int32 `db:"day"`
		Cid int64 `db:"calendar_id"`
		Uid int64 `db:"user_id"`
	}
	err = db.Get(&r, "select day, calendar_id, user_id from entries where id = ?", res.Msg.Id)
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
