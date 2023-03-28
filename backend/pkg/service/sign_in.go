package service

import (
	"context"
	"database/sql"

	adventarv1 "github.com/adventar/adventar/backend/pkg/gen/proto/adventar/v1"
	"github.com/adventar/adventar/backend/pkg/model"
	"github.com/adventar/adventar/backend/pkg/util"
	"github.com/bufbuild/connect-go"
	"github.com/m-mizutani/goerr"
)

// SignIn validates the id token.
func (s *Service) SignIn(
	ctx context.Context,
	req *connect.Request[adventarv1.SignInRequest],
) (*connect.Response[adventarv1.User], error) {
	authResult, err := s.verifier.VerifyIDToken(req.Msg.GetJwt())
	if err != nil {
		return nil, goerr.Wrap(err, "Failed to verify token")
	}

	user, err := s.findOrCreateUser(authResult, req.Msg.GetIconUrl())
	if err != nil {
		return nil, goerr.Wrap(err, "Failed find or create user")
	}

	return connect.NewResponse(&adventarv1.User{
		Id:      user.ID,
		Name:    user.Name,
		IconUrl: user.IconURL,
	}), nil
}

func (s *Service) findOrCreateUser(authResult *util.AuthResult, iconURL string) (*model.User, error) {
	var userID int64
	err := s.db.Get(&userID, "select id from users where auth_provider = ? and auth_uid = ?", authResult.AuthProvider, authResult.AuthUID)

	if err != nil && err != sql.ErrNoRows {
		return nil, goerr.Wrap(err, "Failed query to fetch user")
	}

	if iconURL == "" {
		iconURL = authResult.IconURL
	}

	if err == sql.ErrNoRows {
		res, err := s.db.Exec(
			"insert into users (name, auth_uid, auth_provider, icon_url) values (?, ?, ?, ?)",
			authResult.Name, authResult.AuthUID, authResult.AuthProvider, iconURL,
		)
		if err != nil {
			return nil, goerr.Wrap(err, "Failed query to insert into user")
		}

		userID, err = res.LastInsertId()
		if err != nil {
			return nil, goerr.Wrap(err, "Failed to get last id")
		}
	} else {
		_, err := s.db.Exec("update users set icon_url = ? where id = ?", iconURL, userID)
		if err != nil {
			return nil, goerr.Wrap(err, "Failed query to update user")
		}
	}

	var user model.User
	err = s.db.Get(&user, "select id, name, icon_url from users where id = ?", userID)
	if err != nil {
		return nil, goerr.Wrap(err, "Failed query to fetch user")
	}

	return &user, nil
}
