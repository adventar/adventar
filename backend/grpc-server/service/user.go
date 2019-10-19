package service

import (
	"context"
	"database/sql"

	"golang.org/x/xerrors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/adventar/adventar/backend/grpc-server/grpc/adventar/v1"
	"github.com/adventar/adventar/backend/grpc-server/model"
)

// SignIn validates the id token.
func (s *Service) SignIn(ctx context.Context, in *pb.SignInRequest) (*pb.User, error) {
	authResult, err := s.verifier.VerifyIDToken(in.GetJwt())
	if err != nil {
		return nil, xerrors.Errorf("Failed to verify token: %w", err)
	}

	var userID int64
	var userName string
	err = s.db.QueryRow("select id, name from users where auth_provider = ? and auth_uid = ?", authResult.AuthProvider, authResult.AuthUID).Scan(&userID, &userName)

	if err != nil && err != sql.ErrNoRows {
		return nil, xerrors.Errorf("Failed query to fetch user: %w", err)
	}

	u := &pb.User{Name: authResult.Name, IconUrl: authResult.IconURL}

	if err == sql.ErrNoRows {
		stmt, err := s.db.Prepare("insert into users (name, auth_uid, auth_provider, icon_url) values (?, ?, ?, ?)")
		if err != nil {
			return nil, xerrors.Errorf("Failed to prepare query: %w", err)
		}
		defer stmt.Close()
		res, err := stmt.Exec(authResult.Name, authResult.AuthUID, authResult.AuthProvider, authResult.IconURL)
		if err != nil {
			return nil, xerrors.Errorf("Failed query to insert into user: %w", err)
		}
		u.Id, err = res.LastInsertId()
		if err != nil {
			return nil, xerrors.Errorf("Failed to get last id: %w", err)
		}
	} else {
		stmt, err := s.db.Prepare("update users set icon_url = ? where id = ?")
		if err != nil {
			return nil, xerrors.Errorf("Failed to prepare query: %w", err)
		}
		defer stmt.Close()
		_, err = stmt.Exec(authResult.IconURL, userID)
		if err != nil {
			return nil, xerrors.Errorf("Failed query to update user: %w", err)
		}
		u.Id = userID
		u.Name = userName
	}

	return u, nil
}

// GetUser returns a user info.
func (s *Service) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.User, error) {
	var user model.User
	row := s.db.QueryRow("select id, name, icon_url from users where id = ?", in.GetUserId())
	err := row.Scan(&user.ID, &user.Name, &user.IconURL)

	if err == sql.ErrNoRows {
		return nil, status.Error(codes.NotFound, "User not found")
	}

	if err != nil {
		return nil, xerrors.Errorf("Failed query to fetch user: %w", err)
	}

	return &pb.User{Id: user.ID, Name: user.Name, IconUrl: user.IconURL}, nil
}

// UpdateUser updates user info.
func (s *Service) UpdateUser(ctx context.Context, in *pb.UpdateUserRequest) (*pb.User, error) {
	currentUser, err := s.getCurrentUser(ctx)
	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, "Invalid token")
	}
	name := in.GetName()
	if name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Name is blank")
	}

	stmt, err := s.db.Prepare("update users set name = ? where id = ?")
	if err != nil {
		return nil, xerrors.Errorf("Failed to prepare query: %w", err)
	}

	_, err = stmt.Exec(name, currentUser.ID)
	if err != nil {
		return nil, xerrors.Errorf("Failed query to update user: %w", err)
	}

	return &pb.User{Id: currentUser.ID, Name: name, IconUrl: currentUser.IconURL}, nil
}
