package service

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	pb "github.com/adventar/adventar/api-server/grpc-server/grpc/adventar/v1"
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

	stmt, err := s.db.Prepare("update entries set comment = ?, url = ? where id = ? and user_id = ?")
	if err != nil {
		return nil, xerrors.Errorf("Failed to prepare query: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(in.GetComment(), inURL, in.GetEntryId(), currentUser.ID)
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
		stmt, err = s.db.Prepare("update entries set title = ?, image_url = ? where id = ? and user_id = ?")
		if err != nil {
			return nil, xerrors.Errorf("Failed to prepare query: %w", err)
		}
		defer stmt.Close()

		_, err = stmt.Exec(title, imageURL, in.GetEntryId(), currentUser.ID)
		if err != nil {
			return nil, xerrors.Errorf("Failed query to update entry: %w", err)
		}
	}

	var comment string
	var url string
	var title string
	var imageURL string
	err = s.db.QueryRow("select comment, url, title, image_url from entries where id = ?", in.GetEntryId()).Scan(&comment, &url, &title, &imageURL)
	if err == sql.ErrNoRows {
		return nil, status.Errorf(codes.NotFound, "Entry not found")
	}
	if err != nil {
		return nil, xerrors.Errorf("Failed query to fetch entry: %w", err)
	}

	return &pb.Entry{Id: in.GetEntryId(), Comment: comment, Url: url, Title: title, ImageUrl: convertImageURL(imageURL)}, nil
}
