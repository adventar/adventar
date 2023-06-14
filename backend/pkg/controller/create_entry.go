package controller

import (
	"context"

	adventarv1 "github.com/adventar/adventar/backend/pkg/gen/proto/adventar/v1"
	"github.com/adventar/adventar/backend/pkg/usecase"
	"github.com/bufbuild/connect-go"
	"github.com/m-mizutani/goerr"
)

// CreateEntry creates a entry.
func (x *Controller) CreateEntry(
	ctx context.Context,
	req *connect.Request[adventarv1.CreateEntryRequest],
) (*connect.Response[adventarv1.Entry], error) {
	currentUser, err := x.authenticate(ctx)

	if err != nil {
		return nil, goerr.Wrap(err, "Failed to authenticate")
	}

	entry, err := x.usecase.CreateEntry(&usecase.CreateEntryInput{
		CalendarID: req.Msg.CalendarId,
		UserID:     currentUser.ID,
		Day:        req.Msg.Day,
	})

	if err != nil {
		return nil, goerr.Wrap(err, "Failed to create entry")
	}

	return connect.NewResponse(entry.ToProto()), nil
}
