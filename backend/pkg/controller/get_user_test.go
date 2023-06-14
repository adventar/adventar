package controller_test

import (
	"context"
	"testing"

	adventarv1 "github.com/adventar/adventar/backend/pkg/gen/proto/adventar/v1"
	"github.com/bufbuild/connect-go"
)

func TestGetUser(t *testing.T) {
	cleanupDatabase()

	u := &user{name: "foo", authUID: "xxx", authProvider: "google"}
	createUser(t, u)

	req := connect.NewRequest(&adventarv1.GetUserRequest{UserId: u.id})
	ctx := context.Background()

	res, err := service.GetUser(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	if res.Msg.Id != req.Msg.UserId {
		t.Errorf("actual: %d, expected: %d", res.Msg.Id, req.Msg.UserId)
	}
}
