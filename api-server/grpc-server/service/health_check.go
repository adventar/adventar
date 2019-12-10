package service

import (
	"context"

	pb "github.com/adventar/adventar/api-server/grpc-server/grpc/adventar/v1"
)

// HealthCheck returns current status.
func (s *Service) HealthCheck(ctx context.Context, in *pb.HealthCheckRequest) (*pb.HealthCheckResponse, error) {
	return &pb.HealthCheckResponse{State: "healthy"}, nil
}
