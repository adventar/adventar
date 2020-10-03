package service_test

import (
	"context"
	"testing"

	pb "github.com/adventar/adventar/api-server/grpc-server/grpc/adventar/v1"
	"google.golang.org/grpc/metadata"
)

func TestUpdateUser(t *testing.T) {
	cleanupDatabase()

	u := &user{name: "foo", authUID: "xxx", authProvider: "google"}
	createUser(t, u)

	in := &pb.UpdateUserRequest{Name: "changed"}
	md := make(map[string][]string)
	md["authorization"] = append(md["authorization"], u.authUID)
	ctx := metadata.NewIncomingContext(context.Background(), md)

	_, err := service.UpdateUser(ctx, in)
	if err != nil {
		t.Fatal(err)
	}

	var name string
	err = db.QueryRow("select name from users where id = ?", u.id).Scan(&name)
	if err != nil {
		t.Fatal(err)
	}

	if name != "changed" {
		t.Errorf("actual: %s, expected: %s", name, "changed")
	}
}
