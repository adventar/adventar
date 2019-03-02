package main

import (
	"context"
	"log"
	"net"

	pb "github.com/adventar/adventar/grpc-server/adventar/v1"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type server struct{}

func (s *server) GetCalendar(ctx context.Context, in *pb.GetCalendarRequest) (*pb.Calendar, error) {
	return &pb.Calendar{Id: in.GetId()}, nil
}

func main() {
	log.Printf("listening at %s\n", port)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln(err)
	}

	s := grpc.NewServer()
	pb.RegisterAdventarServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
