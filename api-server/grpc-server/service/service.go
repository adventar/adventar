package service

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/bugsnag/bugsnag-go"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/adventar/adventar/api-server/grpc-server/grpc/adventar/v1"
	"github.com/adventar/adventar/api-server/grpc-server/util"
)

type verifier interface {
	VerifyIDToken(string) (*util.AuthResult, error)
}

type metaFetcher interface {
	Fetch(string) (*util.SiteMeta, error)
}

// Service holds data used by grpc functions.
type Service struct {
	db          *sql.DB
	verifier    verifier
	metaFetcher metaFetcher
}

// NewService creates a new Service.
func NewService(db *sql.DB, verifier verifier, metaFetcher metaFetcher) *Service {
	return &Service{db: db, verifier: verifier, metaFetcher: metaFetcher}
}

// Serve serves the service
func (s *Service) Serve(addr string) {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetOutput(os.Stdout)
	// logrus.SetFormatter(&logrus.JSONFormatter{})
	logger := logrus.WithFields(logrus.Fields{})
	opts := []grpc_logrus.Option{
		grpc_logrus.WithDurationField(func(duration time.Duration) (key string, value interface{}) {
			return "grpc.time_ns", duration.Nanoseconds()
		}),
	}

	grpc_logrus.ReplaceGrpcLogger(logger)

	bugsnagAPIKey := os.Getenv("BUGSNAG_API_KEY")
	if bugsnagAPIKey != "" {
		bugsnag.Configure(bugsnag.Configuration{
			APIKey: bugsnagAPIKey,
		})
	}

	server := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			grpc_logrus.UnaryServerInterceptor(logger, opts...),
			func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (_ interface{}, err error) {
				defer func() {
					if r := recover(); r != nil {
						err = grpc.Errorf(codes.Internal, "Internal Server Error")
						fmt.Printf("%s\n", r)
						if bugsnagAPIKey != "" {
							bugsnag.Notify(fmt.Errorf("%s", r), ctx)
						}
					}
				}()
				resp, err := handler(ctx, req)
				s, _ := status.FromError(err)
				if s.Code() == codes.Unknown {
					stacktrace := fmt.Sprintf("%+v\n", err)
					fmt.Print(stacktrace)
					if bugsnagAPIKey != "" {
						bugsnag.Notify(err, ctx, bugsnag.MetaData{"info": {"stacktrace": stacktrace}})
					}
					err = grpc.Errorf(codes.Internal, "Internal Server Error")
				}
				return resp, err
			},
		),
	)
	pb.RegisterAdventarServer(server, s)
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
