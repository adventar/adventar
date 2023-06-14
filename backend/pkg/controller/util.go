package controller

import (
	"context"
	"database/sql"

	"github.com/adventar/adventar/backend/pkg/domain/model"
	"github.com/adventar/adventar/backend/pkg/domain/types"
	"github.com/adventar/adventar/backend/pkg/gen/sqlc/adventar_db"
	"github.com/m-mizutani/goerr"
)

func (x *Service) authenticate(ctx context.Context) (*model.User, error) {
	metadata, err := GetRequestMetadata(ctx)
	if err != nil {
		return nil, goerr.Wrap(err, "Failed to get request metadata")
	}

	if metadata.authToken == "" {
		return nil, goerr.Wrap(types.ErrPermissionDenied, "Authentication failed").With("auth_token", "empty")
	}

	authResult, err := x.verifier.VerifyIDToken(metadata.authToken)
	if err != nil {
		return nil, goerr.Wrap(types.ErrPermissionDenied, "Failed to verify token")
	}

	user, err := x.usecase.GetUserByAuthInfo(authResult.AuthProvider, authResult.AuthUID)
	if err == sql.ErrNoRows {
		return nil, goerr.Wrap(types.ErrPermissionDenied, "Authentication failed").With("auth_result", authResult)
	}

	if err != nil {
		return nil, goerr.Wrap(err, "Failed query to fetch user")
	}

	return user, nil
}

func (s *Service) queries() *adventar_db.Queries {
	return s.clients.DB().Queries()
}
