package service

import (
	"context"
	"database/sql"

	pb "github.com/adventar/adventar/api-server/grpc-server/grpc/adventar/v1"
	"github.com/adventar/adventar/api-server/grpc-server/model"
	"github.com/adventar/adventar/api-server/grpc-server/util"
	"golang.org/x/xerrors"
)

// SignIn validates the id token.
func (s *Service) SignIn(ctx context.Context, in *pb.SignInRequest) (*pb.User, error) {
	authResult, err := s.verifier.VerifyIDToken(in.GetJwt())
	if err != nil {
		return nil, xerrors.Errorf("Failed to verify token: %w", err)
	}

	user, err := s.findOrCreateUser(authResult, in.GetIconUrl())
	if err != nil {
		return nil, xerrors.Errorf("Failed find or create user: %w", err)
	}

	return &pb.User{Id: user.ID, Name: user.Name, IconUrl: user.IconURL}, nil
}

func (s *Service) findOrCreateUser(authResult *util.AuthResult, iconURL string) (*model.User, error) {
	var userID int64
	err := s.db.Get(&userID, "select id from users where auth_provider = ? and auth_uid = ?", authResult.AuthProvider, authResult.AuthUID)

	if err != nil && err != sql.ErrNoRows {
		return nil, xerrors.Errorf("Failed query to fetch user: %w", err)
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
			return nil, xerrors.Errorf("Failed query to insert into user: %w", err)
		}

		userID, err = res.LastInsertId()
		if err != nil {
			return nil, xerrors.Errorf("Failed to get last id: %w", err)
		}
	} else {
		_, err := s.db.Exec("update users set icon_url = ? where id = ?", iconURL, userID)
		if err != nil {
			return nil, xerrors.Errorf("Failed query to update user: %w", err)
		}
	}

	var user model.User
	err = s.db.Get(&user, "select id, name, icon_url from users where id = ?", userID)
	if err != nil {
		return nil, xerrors.Errorf("Failed query to fetch user: %w", err)
	}

	return &user, nil
}
