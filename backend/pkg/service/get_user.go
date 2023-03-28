package service

import (
	"context"
	"database/sql"
	"errors"

	adventarv1 "github.com/adventar/adventar/backend/pkg/gen/proto/adventar/v1"
	"github.com/bufbuild/connect-go"
	"github.com/m-mizutani/goerr"
)

// GetUser returns a user info.
func (s *Service) GetUser(
	ctx context.Context,
	req *connect.Request[adventarv1.GetUserRequest],
) (*connect.Response[adventarv1.User], error) {
	userId := req.Msg.GetUserId()
	user, err := s.clients.DB().Queries().GetUserById(context.Background(), userId)

	if err == sql.ErrNoRows {
		return nil, connect.NewError(connect.CodeNotFound, errors.New("User not found"))
	}

	if err != nil {
		return nil, goerr.Wrap(err, "Failed query to fetch user")
	}

	return connect.NewResponse(&adventarv1.User{
		Id:      user.ID,
		Name:    user.Name,
		IconUrl: user.IconUrl,
	}), nil
}
