package controller

import (
	"context"

	adventarv1 "github.com/adventar/adventar/backend/pkg/gen/proto/adventar/v1"
	"github.com/bufbuild/connect-go"
	"github.com/m-mizutani/goerr"
)

// SignIn validates the id token.
func (x *Service) SignIn(
	ctx context.Context,
	req *connect.Request[adventarv1.SignInRequest],
) (*connect.Response[adventarv1.User], error) {
	authResult, err := x.verifier.VerifyIDToken(req.Msg.GetJwt())
	if err != nil {
		return nil, goerr.Wrap(err, "Failed to verify token")
	}

	user, err := x.usecase.FindOrCreateUser(authResult, req.Msg.GetIconUrl())
	if err != nil {
		return nil, goerr.Wrap(err, "Failed find or create user")
	}

	return connect.NewResponse(user.ToProto()), nil
}
