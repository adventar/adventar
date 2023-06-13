package service

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/adventar/adventar/backend/pkg/domain/types"
	"github.com/adventar/adventar/backend/pkg/gen/proto/adventar/v1/adventarv1connect"
	"github.com/adventar/adventar/backend/pkg/infra"
	"github.com/adventar/adventar/backend/pkg/usecase"
	"github.com/adventar/adventar/backend/pkg/util"
	"github.com/bufbuild/connect-go"
	"github.com/getsentry/sentry-go"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/m-mizutani/goerr"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type verifier interface {
	VerifyIDToken(string) (*util.AuthResult, error)
}

// Service holds data used by grpc functions.
type Service struct {
	db       *sqlx.DB
	verifier verifier
	usecase  *usecase.Usecase
	clients  *infra.Clients
}

// NewService creates a new Service.
func NewService(db *sqlx.DB, verifier verifier, usecase *usecase.Usecase, clients *infra.Clients) *Service {
	return &Service{
		db:       db,
		verifier: verifier,
		usecase:  usecase,
		clients:  clients,
	}
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
	if err != nil {
		panic(err)
	}
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
					logger := util.Logger.Error()
					var goErr *goerr.Error
					if errors.As(err, &goErr) {
						for k, v := range goErr.Values() {
							logger = logger.Any(fmt.Sprintf("error.%v", k), v)
						}
					}
					logger.
						Str("procedure", procedure).
						Str("duration", duration).
						Str("code", connect.CodeOf(err).String()).
						Err(err).
						Msg("finished unary call")
				} else {
					util.Logger.Info().
						Str("procedure", procedure).
						Str("duration", duration).
						Msg("finished unary call")
				}

				return response, err
			})
		},
	)

	sentryDsn := os.Getenv("SENTRY_DSN")
	if sentryDsn != "" {
		err := sentry.Init(sentry.ClientOptions{Dsn: sentryDsn})

		if err != nil {
			panic("Sentry initialize failed: " + err.Error())
		}
	}

	errorHandlerInterceptor := connect.UnaryInterceptorFunc(func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(ctx context.Context, request connect.AnyRequest) (connect.AnyResponse, error) {
			response, err := next(ctx, request)
			if err == nil {
				return response, nil
			}

			switch {
			case errors.Is(err, types.ErrRecordNotFound):
				err = connect.NewError(connect.CodeNotFound, err)
			}

			if connect.CodeOf(err) == connect.CodeUnknown {
				tags := map[string]interface{}{
					"procedure": request.Spec().Procedure,
				}
				var goErr *goerr.Error
				if errors.As(err, &goErr) {
					for k, v := range goErr.Values() {
						tags[k] = v
					}
				}
				if sentryDsn != "" {
					sentry.ConfigureScope(func(scope *sentry.Scope) {
						for k, v := range tags {
							scope.SetTag(k, fmt.Sprintf("%v", v))
						}
					})
					sentry.CaptureException(err)
				}
			}
			printErrorStacks(err)
			return response, err
		})
	})

	metadataInterceptor := connect.UnaryInterceptorFunc(func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(ctx context.Context, request connect.AnyRequest) (connect.AnyResponse, error) {
			ctx = SetRequestMetadata(ctx, request)
			return next(ctx, request)
		})
	})

	return connect.WithInterceptors(metadataInterceptor, loggingInterceptor, errorHandlerInterceptor)
}

func printErrorStacks(err error) {
	var s []string
	for err != nil {
		var e *goerr.Error
		if errors.As(err, &e) {
			s = append(s, fmt.Sprintf("%+v", e.StackTrace()[0]))
		}
		err = errors.Unwrap(err)
	}
	if len(s) > 0 {
		util.Logger.Debug().Msg("Error StackTrace:\n" + strings.Join(s, "\n"))
	}
}
