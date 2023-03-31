package service

import (
	"context"

	adventarv1 "github.com/adventar/adventar/backend/pkg/gen/proto/adventar/v1"
	"github.com/adventar/adventar/backend/pkg/usecase"
	"github.com/bufbuild/connect-go"
)

// CreateEntry creates a entry.
func (x *Service) CreateEntry(
	ctx context.Context,
	req *connect.Request[adventarv1.CreateEntryRequest],
) (*connect.Response[adventarv1.Entry], error) {
	currentUser, err := x.authenticate(ctx)

	if err != nil {
		return nil, err
	}

	entry, err := x.usecase.CreateEntry(&usecase.CreateEntryInput{
		CalendarID: req.Msg.CalendarId,
		UserID:     currentUser.ID,
		Day:        req.Msg.Day,
	})

	if err != nil {
		return nil, err
	}

	return connect.NewResponse(entry.ToProto()), nil
}
