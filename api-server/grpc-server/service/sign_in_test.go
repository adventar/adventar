package service_test

import (
	"context"
	"testing"

	pb "github.com/adventar/adventar/api-server/grpc-server/grpc/adventar/v1"
)

func TestSignInIfUserExists(t *testing.T) {
	cleanupDatabase()

	u := &user{name: "foo", authUID: "xxx", authProvider: "google"}
	createUser(t, u)

	in := &pb.SignInRequest{Jwt: u.authUID, IconUrl: "http://xxx/icon"}
	ctx := context.Background()
	_, err := service.SignIn(ctx, in)
	if err != nil {
		t.Fatal(err)
	}
	var iconURL string
	err = db.Get(&iconURL, "select icon_url from users")
	if err != nil {
		t.Fatal(err)
	}
	expected := "http://xxx/icon"
	if iconURL != expected {
		t.Errorf("actual: %s, expected: %s", iconURL, expected)
	}
}

func TestSignInIfUserDoesNotExist(t *testing.T) {
	cleanupDatabase()
	in := &pb.SignInRequest{Jwt: "", IconUrl: "http://xxx/icon"}
	ctx := context.Background()
	out, err := service.SignIn(ctx, in)
	if err != nil {
		t.Fatal(err)
	}
	var count int
	err = db.Get(&count, "select count(*) users")
	if err != nil {
		t.Fatal(err)
	}
	if count != 1 {
		t.Errorf("actual: %d, expected: 1", count)
	}
	expected := "http://xxx/icon"
	if out.IconUrl != expected {
		t.Errorf("actual: %s, expected: %s", out.IconUrl, expected)
	}
}
