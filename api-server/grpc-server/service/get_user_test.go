package service_test

import (
	"context"
	"testing"

	pb "github.com/adventar/adventar/api-server/grpc-server/grpc/adventar/v1"
)

func TestGetUser(t *testing.T) {
	cleanupDatabase()

	u := &user{name: "foo", authUID: "xxx", authProvider: "google"}
	createUser(t, u)

	in := &pb.GetUserRequest{UserId: u.id}
	ctx := context.Background()

	res, err := service.GetUser(ctx, in)
	if err != nil {
		t.Fatal(err)
	}

	if res.Id != in.UserId {
		t.Errorf("actual: %d, expected: %d", res.Id, in.UserId)
	}
}
