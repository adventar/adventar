package usecase

import (
	"context"
	"database/sql"

	"github.com/adventar/adventar/backend/pkg/domain/model"
	"github.com/adventar/adventar/backend/pkg/domain/types"
	"github.com/adventar/adventar/backend/pkg/gen/sqlc/adventar_db"
	"github.com/m-mizutani/goerr"
)

func (x *Usecase) GetUserById(id int64) (*model.User, error) {
	user, err := x.queries.GetUserById(context.Background(), id)

	if err == sql.ErrNoRows {
		return nil, goerr.Wrap(types.ErrRecordNotFound, "User not found").With("user_id", id)
	}

	if err != nil {
		return nil, goerr.Wrap(err, "Failed query to fetch user")
	}

	return &model.User{
		ID:      user.ID,
		Name:    user.Name,
		IconURL: user.IconUrl,
	}, nil
}

func (x *Usecase) GetUserByAuthInfo(provider string, uid string) (*model.User, error) {
	user, err := x.queries.GetUserByAuthInfo(context.Background(), adventar_db.GetUserByAuthInfoParams{AuthProvider: provider, AuthUid: uid})

	if err == sql.ErrNoRows {
		return nil, goerr.Wrap(types.ErrRecordNotFound, "User not found").
			With("auth_provider", provider).
			With("auth_uid", uid)
	}

	if err != nil {
		return nil, goerr.Wrap(err, "Failed query to fetch user")
	}

	return &model.User{
		ID:      user.ID,
		Name:    user.Name,
		IconURL: user.IconUrl,
	}, nil
}

type UpdateUserInput struct {
	UserID int64
	Name   string
}

func (x *Usecase) UpdateUser(input *UpdateUserInput) error {
	if input.Name == "" {
		return goerr.Wrap(types.ErrInvalidArgument, "Name is required")
	}

	err := x.queries.UpdateUser(context.Background(), adventar_db.UpdateUserParams{
		ID:   input.UserID,
		Name: input.Name,
	})

	if err != nil {
		return goerr.Wrap(err, "Failed query to update user")
	}

	return nil
}
