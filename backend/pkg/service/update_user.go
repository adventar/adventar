package service

import (
	"context"

	adventarv1 "github.com/adventar/adventar/backend/pkg/gen/proto/adventar/v1"
	"github.com/adventar/adventar/backend/pkg/usecase"
	"github.com/bufbuild/connect-go"
	"github.com/m-mizutani/goerr"
)

// UpdateUser updates user info.
func (x *Service) UpdateUser(
	ctx context.Context,
	req *connect.Request[adventarv1.UpdateUserRequest],
) (*connect.Response[adventarv1.User], error) {
	currentUser, err := x.authenticate(ctx)

	if err != nil {
		return nil, goerr.Wrap(err, "Failed to authenticate")
	}

	err = x.usecase.UpdateUser(&usecase.UpdateUserInput{
		UserID: currentUser.ID,
		Name:   req.Msg.Name,
	})

	if err != nil {
		return nil, goerr.Wrap(err, "Failed to update user")
	}

	return connect.NewResponse(&adventarv1.User{
		Id:      currentUser.ID,
		Name:    req.Msg.Name,
		IconUrl: currentUser.IconURL,
	}), nil
}
