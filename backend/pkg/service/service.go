package service

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/adventar/adventar/backend/pkg/gen/adventar/v1/adventarv1connect"
	"github.com/adventar/adventar/backend/pkg/util"
	"github.com/bufbuild/connect-go"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type verifier interface {
	VerifyIDToken(string) (*util.AuthResult, error)
}

type metaFetcher interface {
	Fetch(string) (*util.SiteMeta, error)
}

// Service holds data used by grpc functions.
type Service struct {
	db          *sqlx.DB
	verifier    verifier
	metaFetcher metaFetcher
}

// NewService creates a new Service.
func NewService(db *sqlx.DB, verifier verifier, metaFetcher metaFetcher) *Service {
	return &Service{db: db, verifier: verifier, metaFetcher: metaFetcher}
}

// Serve serves the service
func (s *Service) Serve(addr string) {
	mux := http.NewServeMux()
	interceptors := createInterceptors()
	mux.Handle(adventarv1connect.NewAdventarHandler(s, interceptors))
	mux.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "ok\n")
	})
	log.Println("Listen: 8080")
	err := http.ListenAndServe(
		":8080",
		h2c.NewHandler(mux, &http2.Server{}),
	)
	log.Fatalf("listen failed: %v", err)
}

func createInterceptors() connect.HandlerOption {
	// logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)
	loggingInterceptor := connect.UnaryInterceptorFunc(
		func(next connect.UnaryFunc) connect.UnaryFunc {
			return connect.UnaryFunc(func(ctx context.Context, request connect.AnyRequest) (connect.AnyResponse, error) {
				logger := logrus.WithField("request", request.Spec().Procedure)
				logger.Info("started unary call")
				start := time.Now()
				response, err := next(ctx, request)
				logger = logger.WithField("time", fmt.Sprintf("%dms", time.Since(start).Milliseconds()))
				if err != nil {
					logger.WithFields(logrus.Fields{
						"code":  connect.CodeOf(err),
						"error": err,
					}).Error("finished unary call")
				} else {
					logger.Info("finished unary call")
				}
				return response, err
			})
		},
	)

	return connect.WithInterceptors(loggingInterceptor)
}

/*
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
*/
