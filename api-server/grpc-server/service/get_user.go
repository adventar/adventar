package service

import (
	"context"
	"database/sql"

	"golang.org/x/xerrors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/adventar/adventar/api-server/grpc-server/grpc/adventar/v1"
	"github.com/adventar/adventar/api-server/grpc-server/model"
)

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
