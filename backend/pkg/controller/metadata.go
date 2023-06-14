package controller

import (
	"context"
	"time"

	"github.com/bufbuild/connect-go"
	"github.com/m-mizutani/goerr"
)

const requestMetadataKey = "request_metadata"

type RequestMetadata struct {
	requestedAt time.Time
	authToken   string
}

func SetRequestMetadata(ctx context.Context, request connect.AnyRequest) context.Context {
	metadata := &RequestMetadata{
		requestedAt: time.Now(),
		authToken:   request.Header().Get("authorization"),
	}
	return context.WithValue(ctx, requestMetadataKey, metadata)
}

func GetRequestMetadata(ctx context.Context) (*RequestMetadata, error) {
	v := ctx.Value(requestMetadataKey)
	metadata, ok := v.(*RequestMetadata)

	if !ok {
		return nil, goerr.New("Invalid request metadata in context")
	}

	return metadata, nil
}
