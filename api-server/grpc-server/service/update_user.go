package service

import (
	"context"

	pb "github.com/adventar/adventar/api-server/grpc-server/grpc/adventar/v1"
	"golang.org/x/xerrors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// UpdateUser updates user info.
func (s *Service) UpdateUser(ctx context.Context, in *pb.UpdateUserRequest) (*pb.User, error) {
	currentUser, err := s.getCurrentUser(ctx)
	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, "Authentication failed")
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
