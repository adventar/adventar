package service_test

import (
	"context"
	"testing"

	pb "github.com/adventar/adventar/api-server/grpc-server/grpc/adventar/v1"
)

func TestHealthCheck(t *testing.T) {
	in := &pb.HealthCheckRequest{}

	ctx := context.Background()

	res, err := service.HealthCheck(ctx, in)
	if err != nil {
		t.Fatal(err)
	}

	if res.State != "healthy" {
		t.Errorf("actual: %s, expected: %s", res.State, "healthy")
	}
}
