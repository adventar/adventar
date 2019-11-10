package service

import (
	"context"
	"crypto/sha1"
	"fmt"
	"net/url"
	"os"
	"strings"

	"golang.org/x/xerrors"
	"google.golang.org/grpc/metadata"

	pb "github.com/adventar/adventar/backend/grpc-server/grpc/adventar/v1"
	"github.com/adventar/adventar/backend/grpc-server/model"
)

func (s *Service) getCurrentUser(ctx context.Context) (*model.User, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, xerrors.Errorf("Metadata not found")
	}

	values := md["authorization"]
	if len(values) == 0 {
		return nil, xerrors.Errorf("Authorization metadata not found")
	}

	authResult, err := s.verifier.VerifyIDToken(values[0])
	if err != nil {
		return nil, xerrors.Errorf("Failed to verify token: %w", err)
	}

	var user model.User
	err = s.db.QueryRow("select id, name, icon_url from users where auth_provider = ? and auth_uid = ?", authResult.AuthProvider, authResult.AuthUID).Scan(&user.ID, &user.Name, &user.IconURL)
	if err != nil {
		return nil, xerrors.Errorf("Failed query to fetch user: %w", err)
	}

	return &user, nil
}

func (s *Service) bindEntryCount(calendars []*pb.Calendar) error {
	ids := []interface{}{}
	interpolations := []string{}

	for _, c := range calendars {
		ids = append(ids, c.Id)
		interpolations = append(interpolations, "?")
	}

	sql := fmt.Sprintf("select calendar_id, count(*) from entries where calendar_id in (%s) group by calendar_id", strings.Join(interpolations, ","))
	rows, err := s.db.Query(sql, ids...)
	if err != nil {
		return xerrors.Errorf("Failed query to fetch calendar: %w", err)
	}

	entryCounts := map[int64]int32{}
	for rows.Next() {
		var cid int64
		var count int32
		if err := rows.Scan(&cid, &count); err != nil {
			return xerrors.Errorf("Failed to scan row: %w", err)
		}
		entryCounts[cid] = count
	}

	for _, c := range calendars {
		c.EntryCount = entryCounts[c.Id]
	}

	return nil
}

func convertImageURL(imageURL string) string {
	endpoint := os.Getenv("IMAGE_SERVER_ENDPOINT")
	if endpoint == "" {
		return imageURL
	}
	salt := os.Getenv("DIGEST_SALT")
	h := sha1.New()
	h.Write([]byte(imageURL + salt))

	return fmt.Sprintf("%s/%x?url=%s", endpoint, h.Sum(nil), url.QueryEscape(imageURL))
}
