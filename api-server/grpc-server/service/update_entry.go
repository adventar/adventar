package service

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	pb "github.com/adventar/adventar/api-server/grpc-server/grpc/adventar/v1"
	"github.com/adventar/adventar/api-server/grpc-server/model"
	"golang.org/x/xerrors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// UpdateEntry updates the entry.
func (s *Service) UpdateEntry(ctx context.Context, in *pb.UpdateEntryRequest) (*pb.Entry, error) {
	currentUser, err := s.getCurrentUser(ctx)
	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, "Authentication failed")
	}

	inURL := strings.TrimSpace(in.GetUrl())
	if inURL != "" && !isValidURL(inURL) {
		return nil, status.Errorf(codes.InvalidArgument, "URL is invalid")
	}

	_, err = s.db.Exec(
		"update entries set comment = ?, url = ? where id = ? and user_id = ?",
		in.GetComment(), inURL, in.GetEntryId(), currentUser.ID,
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
			title, imageURL, in.GetEntryId(), currentUser.ID,
		)
		if err != nil {
			return nil, xerrors.Errorf("Failed query to update entry: %w", err)
		}
	}

	var entry model.Entry
	err = s.db.Get(&entry, "select id, comment, url, title, image_url from entries where id = ?", in.GetEntryId())
	if err == sql.ErrNoRows {
		return nil, status.Errorf(codes.NotFound, "Entry not found")
	}
	if err != nil {
		return nil, xerrors.Errorf("Failed query to fetch entry: %w", err)
	}

	return &pb.Entry{
		Id:       in.GetEntryId(),
		Comment:  entry.Comment,
		Url:      entry.URL,
		Title:    entry.Title,
		ImageUrl: convertImageURL(entry.ImageURL),
	}, nil
}
