package service_test

import (
	"context"
	"testing"

	adventarv1 "github.com/adventar/adventar/backend/pkg/gen/proto/adventar/v1"
	"github.com/bufbuild/connect-go"
)

func TestSignInIfUserExists(t *testing.T) {
	cleanupDatabase()

	u := &user{name: "foo", authUID: "xxx", authProvider: "google"}
	createUser(t, u)

	req := connect.NewRequest(&adventarv1.SignInRequest{Jwt: u.authUID, IconUrl: "http://xxx/icon"})
	ctx := context.Background()
	_, err := service.SignIn(ctx, req)
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
	req := connect.NewRequest(&adventarv1.SignInRequest{Jwt: "", IconUrl: "http://xxx/icon"})
	ctx := context.Background()
	res, err := service.SignIn(ctx, req)
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
	if res.Msg.IconUrl != expected {
		t.Errorf("actual: %s, expected: %s", res.Msg.IconUrl, expected)
	}
}
