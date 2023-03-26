package service

import (
	"context"
	"errors"

	adventarv1 "github.com/adventar/adventar/backend/pkg/gen/adventar/v1"
	"github.com/bufbuild/connect-go"
	"golang.org/x/xerrors"
)

// UpdateUser updates user info.
func (s *Service) UpdateUser(
	ctx context.Context,
	req *connect.Request[adventarv1.UpdateUserRequest],
) (*connect.Response[adventarv1.User], error) {
	currentUser, err := s.getCurrentUser(req.Header())
	if err != nil {
		return nil, connect.NewError(connect.CodePermissionDenied, errors.New("Authentication failed"))
	}

	name := req.Msg.Name
	if name == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("Name is blank"))
	}

	_, err = s.db.Exec("update users set name = ? where id = ?", name, currentUser.ID)
	if err != nil {
		return nil, xerrors.Errorf("Failed query to update user: %w", err)
	}

	return connect.NewResponse(&adventarv1.User{
		Id:      currentUser.ID,
		Name:    name,
		IconUrl: currentUser.IconURL,
	}), nil
}
