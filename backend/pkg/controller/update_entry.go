package controller

import (
	"context"

	adventarv1 "github.com/adventar/adventar/backend/pkg/gen/proto/adventar/v1"
	"github.com/adventar/adventar/backend/pkg/usecase"
	"github.com/bufbuild/connect-go"
	"github.com/m-mizutani/goerr"
)

// UpdateEntry updates the entry.
func (x *Controller) UpdateEntry(
	ctx context.Context,
	req *connect.Request[adventarv1.UpdateEntryRequest],
) (*connect.Response[adventarv1.Entry], error) {
	currentUser, err := x.authenticate(ctx)

	if err != nil {
		return nil, goerr.Wrap(err, "Failed to authenticate")
	}

	entry, err := x.usecase.UpdateEntry(&usecase.UpdateEntryInput{
		EntryID: req.Msg.GetEntryId(),
		Comment: req.Msg.GetComment(),
		URL:     req.Msg.GetUrl(),
		UserID:  currentUser.ID,
	})

	return connect.NewResponse(entry.ToProto()), nil
}
