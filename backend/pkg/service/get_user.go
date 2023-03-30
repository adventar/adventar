package service

import (
	"context"

	adventarv1 "github.com/adventar/adventar/backend/pkg/gen/proto/adventar/v1"
	"github.com/bufbuild/connect-go"
)

// GetUser returns a user info.
func (x *Service) GetUser(
	ctx context.Context,
	req *connect.Request[adventarv1.GetUserRequest],
) (*connect.Response[adventarv1.User], error) {
	userId := req.Msg.GetUserId()
	user, err := x.usecase.GetUserById(userId)

	if err != nil {
		return nil, err
	}

	return connect.NewResponse(user.ToProto()), nil
}
