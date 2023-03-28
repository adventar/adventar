package service

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/adventar/adventar/backend/pkg/gen/adventar/v1/adventarv1connect"
	"github.com/adventar/adventar/backend/pkg/util"
	"github.com/bufbuild/connect-go"
	"github.com/bugsnag/bugsnag-go/v2"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/m-mizutani/goerr"
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
	withRecover := connect.WithRecover(func(_ context.Context, _ connect.Spec, _ http.Header, r any) error {
		return goerr.New("(panic) %v", r)
	})
	mux.Handle(adventarv1connect.NewAdventarHandler(s, interceptors, withRecover))
	mux.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "ok\n")
	})
	util.Logger.Info().Msg("Server started on port 8080")
	err := http.ListenAndServe(
		":8080",
		h2c.NewHandler(mux, &http2.Server{}),
	)
	util.Logger.Fatal().Err(err).Msg("listen failed")
}

func createInterceptors() connect.HandlerOption {
	loggingInterceptor := connect.UnaryInterceptorFunc(
		func(next connect.UnaryFunc) connect.UnaryFunc {
			return connect.UnaryFunc(func(ctx context.Context, request connect.AnyRequest) (connect.AnyResponse, error) {
				procedure := request.Spec().Procedure
				util.Logger.Info().Str("request", procedure).Msg("started unary call")
				start := time.Now()
				response, err := next(ctx, request)
				duration := fmt.Sprintf("%dms", time.Since(start).Milliseconds())

				if err != nil {
					util.Logger.Error().
						Str("request", procedure).
						Str("duration", duration).
						Str("code", connect.CodeOf(err).String()).
						Err(err).
						Msg("finished unary call")
				} else {
					util.Logger.Info().
						Str("request", procedure).
						Str("duration", duration).
						Msg("finished unary call")
				}

				return response, err
			})
		},
	)

	bugsnagAPIKey := os.Getenv("BUGSNAG_API_KEY")
	if bugsnagAPIKey != "" {
		bugsnag.Configure(bugsnag.Configuration{
			APIKey: bugsnagAPIKey,
		})
	}
	errorHandlerInterceptor := connect.UnaryInterceptorFunc(func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(ctx context.Context, request connect.AnyRequest) (connect.AnyResponse, error) {
			response, err := next(ctx, request)
			if err == nil {
				return response, nil
			}
			if connect.CodeOf(err) == connect.CodeUnknown {
				info := map[string]interface{}{}
				stacktrace := fmt.Sprintf("%+v", err)
				util.Logger.Error().Msg(stacktrace)
				info["stacktrace"] = stacktrace
				var goErr *goerr.Error
				if errors.As(err, &goErr) {
					for k, v := range goErr.Values() {
						util.Logger.Error().Any(k, v).Msg("Error context value")
						info[k] = v
					}
				}
				if bugsnagAPIKey != "" {
					bugsnag.Notify(err, bugsnag.MetaData{"info": info})
				}
			}
			return response, err
		})
	})

	return connect.WithInterceptors(loggingInterceptor, errorHandlerInterceptor)
}
