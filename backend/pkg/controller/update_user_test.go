package controller_test

import (
	"context"
	"testing"

	"github.com/adventar/adventar/backend/pkg/controller"
	adventarv1 "github.com/adventar/adventar/backend/pkg/gen/proto/adventar/v1"
	"github.com/bufbuild/connect-go"
)

func TestUpdateUser(t *testing.T) {
	cleanupDatabase()

	u := &user{name: "foo", authUID: "xxx", authProvider: "google"}
	createUser(t, u)

	req := connect.NewRequest(&adventarv1.UpdateUserRequest{Name: "changed"})
	req.Header().Set("authorization", u.authUID)
	ctx := controller.SetRequestMetadata(context.Background(), req)

	_, err := service.UpdateUser(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	var name string
	err = db.Get(&name, "select name from users where id = ?", u.id)
	if err != nil {
		t.Fatal(err)
	}

	if name != "changed" {
		t.Errorf("actual: %s, expected: %s", name, "changed")
	}
}
