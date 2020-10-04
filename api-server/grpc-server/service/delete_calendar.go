package service

import (
	"context"

	pb "github.com/adventar/adventar/api-server/grpc-server/grpc/adventar/v1"
	"github.com/golang/protobuf/ptypes/empty"
	"golang.org/x/xerrors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// DeleteCalendar deletes the calendar.
func (s *Service) DeleteCalendar(ctx context.Context, in *pb.DeleteCalendarRequest) (*empty.Empty, error) {
	currentUser, err := s.getCurrentUser(ctx)
	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, "Authentication failed")
	}

	_, err = s.db.Exec("delete from calendars where id = ? and user_id = ?", in.GetCalendarId(), currentUser.ID)
	if err != nil {
		return nil, xerrors.Errorf("Failed query to delete calendar: %w", err)
	}

	return &empty.Empty{}, nil
}
