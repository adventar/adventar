package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	adventarv1 "github.com/adventar/adventar/backend/pkg/gen/adventar/v1"
	"github.com/adventar/adventar/backend/pkg/model"
	"github.com/bufbuild/connect-go"
	"golang.org/x/xerrors"
)

// UpdateEntry updates the entry.
func (s *Service) UpdateEntry(
	ctx context.Context,
	req *connect.Request[adventarv1.UpdateEntryRequest],
) (*connect.Response[adventarv1.Entry], error) {
	currentUser, err := s.getCurrentUser(req.Header())
	if err != nil {
		return nil, connect.NewError(connect.CodePermissionDenied, errors.New("Authentication failed"))
	}

	inURL := strings.TrimSpace(req.Msg.GetUrl())
	if inURL != "" && !isValidURL(inURL) {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("URL is invalid"))
	}

	_, err = s.db.Exec(
		"update entries set comment = ?, url = ? where id = ? and user_id = ?",
		req.Msg.GetComment(), inURL, req.Msg.GetEntryId(), currentUser.ID,
	)
	if err != nil {
		return nil, xerrors.Errorf("Failed query to update entry: %w", err)
	}

	if inURL != "" {
		m, err := s.metaFetcher.Fetch(inURL)
		var title string
		var imageURL string
		if err != nil {
			fmt.Printf("Failed to fetch url: %s", err)
			title = ""
			imageURL = ""
		} else {
			title = m.Title
			imageURL = m.ImageURL
		}
		_, err = s.db.Exec(
			"update entries set title = ?, image_url = ? where id = ? and user_id = ?",
			title, imageURL, req.Msg.GetEntryId(), currentUser.ID,
		)
		if err != nil {
			return nil, xerrors.Errorf("Failed query to update entry: %w", err)
		}
	}

	var entry model.Entry
	err = s.db.Get(&entry, "select id, comment, url, title, image_url from entries where id = ?", req.Msg.GetEntryId())
	if err == sql.ErrNoRows {
		return nil, connect.NewError(connect.CodeNotFound, errors.New("Entry not found"))
	}
	if err != nil {
		return nil, xerrors.Errorf("Failed query to fetch entry: %w", err)
	}

	return connect.NewResponse(&adventarv1.Entry{
		Id:       req.Msg.GetEntryId(),
		Comment:  entry.Comment,
		Url:      entry.URL,
		Title:    entry.Title,
		ImageUrl: convertImageURL(entry.ImageURL),
	}), nil
}
