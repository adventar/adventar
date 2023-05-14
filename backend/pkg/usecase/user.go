package usecase

import (
	"context"
	"database/sql"

	"github.com/adventar/adventar/backend/pkg/domain/model"
	"github.com/adventar/adventar/backend/pkg/domain/types"
	"github.com/adventar/adventar/backend/pkg/gen/sqlc/adventar_db"
	"github.com/adventar/adventar/backend/pkg/util"
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

func (x *Usecase) FindOrCreateUser(authResult *util.AuthResult, iconURL string) (*model.User, error) {
	user, err := x.queries.GetUserByAuthInfo(context.Background(), adventar_db.GetUserByAuthInfoParams{AuthProvider: authResult.AuthProvider, AuthUid: authResult.AuthUID})

	if err != nil && err != sql.ErrNoRows {
		return nil, goerr.Wrap(err, "Failed query to fetch user")
	}

	if iconURL == "" {
		iconURL = authResult.IconURL
	}

	var userID int64
	if err == sql.ErrNoRows {
		userID, err = x.queries.CreateUser(context.Background(), adventar_db.CreateUserParams{
			Name:         authResult.Name,
			AuthUid:      authResult.AuthUID,
			AuthProvider: authResult.AuthProvider,
			IconUrl:      iconURL,
		})
		if err != nil {
			return nil, goerr.Wrap(err, "Failed query to insert into user")
		}
	} else {
		userID = user.ID
		err := x.queries.UpdateUserIconUrl(context.Background(), adventar_db.UpdateUserIconUrlParams{
			ID:      userID,
			IconUrl: iconURL,
		})

		if err != nil {
			return nil, goerr.Wrap(err, "Failed query to update user")
		}
	}

	user, err = x.queries.GetUserById(context.Background(), userID)

	return &model.User{
		ID:      user.ID,
		Name:    user.Name,
		IconURL: user.IconUrl,
	}, nil
}

type UpdateUserNameInput struct {
	UserID int64
	Name   string
}

func (x *Usecase) UpdateUserName(input *UpdateUserNameInput) error {
	if input.Name == "" {
		return goerr.Wrap(types.ErrInvalidArgument, "Name is required")
	}

	err := x.queries.UpdateUserName(context.Background(), adventar_db.UpdateUserNameParams{
		ID:   input.UserID,
		Name: input.Name,
	})

	if err != nil {
		return goerr.Wrap(err, "Failed query to update user")
	}

	return nil
}
