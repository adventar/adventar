package controller

import (
	"context"

	"github.com/adventar/adventar/backend/pkg/domain/model"
	adventarv1 "github.com/adventar/adventar/backend/pkg/gen/proto/adventar/v1"
	"github.com/bufbuild/connect-go"
	"github.com/m-mizutani/goerr"
	"github.com/m-mizutani/gots/slice"
)

// ListEntries lists entries.
func (x *Controller) ListEntries(
	ctx context.Context,
	req *connect.Request[adventarv1.ListEntriesRequest],
) (*connect.Response[adventarv1.ListEntriesResponse], error) {
	userId := req.Msg.GetUserId()
	year := req.Msg.GetYear()

	var entries []*model.Entry
	var err error
	if req.Msg.GetYear() != 0 {
		entries, err = x.usecase.ListUserEntriesByYear(userId, year)
	} else {
		entries, err = x.usecase.ListUserEntries(userId)
	}

	if err != nil {
		return nil, goerr.Wrap(err, "Failed to list entries")
	}

	return connect.NewResponse(&adventarv1.ListEntriesResponse{
		Entries: slice.Map(entries, func(entry *model.Entry) *adventarv1.Entry {
			return entry.ToProto()
		}),
	}), nil
}
