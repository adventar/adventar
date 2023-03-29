package service

import (
	"crypto/sha1"
	"fmt"
	"net/http"
	"net/url"
	"os"

	sq "github.com/Masterminds/squirrel"
	pb "github.com/adventar/adventar/backend/pkg/gen/proto/adventar/v1"
	adventar_db "github.com/adventar/adventar/backend/pkg/gen/sqlc"
	"github.com/adventar/adventar/backend/pkg/model"
	"github.com/m-mizutani/goerr"
)

func (s *Service) getCurrentUser(header http.Header) (*model.User, error) {
	token := header.Get("authorization")

	authResult, err := s.verifier.VerifyIDToken(token)
	if err != nil {
		return nil, goerr.Wrap(err, "Failed to verify token")
	}

	var user model.User
	err = s.db.Get(&user, "select id, name, icon_url from users where auth_provider = ? and auth_uid = ?", authResult.AuthProvider, authResult.AuthUID)
	if err != nil {
		return nil, goerr.Wrap(err, "Failed query to fetch user")
	}

	return &user, nil
}

func convertImageURL(imageURL string) string {
	endpoint := os.Getenv("IMAGE_SERVER_ENDPOINT")
	if endpoint == "" || imageURL == "" {
		return imageURL
	}
	salt := os.Getenv("IMAGE_DIGEST_SALT")
	h := sha1.New()
	h.Write([]byte(imageURL + salt))

	return fmt.Sprintf("%s/%x?url=%s", endpoint, h.Sum(nil), url.QueryEscape(imageURL))
}

func (s *Service) findEntries(cid int64) ([]*pb.Entry, error) {
	query, args, err := sq.
		Select(makeSelectValue(map[string][]string{
			"entries": {"id", "day", "title", "comment", "url", "image_url"},
			"users":   {"id", "name", "icon_url"},
		})...).
		From("entries").
		Join("users on users.id = entries.user_id").
		Where(sq.Eq{"entries.calendar_id": cid}).
		OrderBy("entries.day").
		ToSql()

	if err != nil {
		return nil, goerr.Wrap(err, "Failed query to create sql")
	}

	rows := []struct {
		Entry model.Entry `db:"entries"`
		User  model.User  `db:"users"`
	}{}

	err = s.db.Select(&rows, query, args...)
	if err != nil {
		return nil, goerr.Wrap(err, "Failed query to fetch entries")
	}

	entries := []*pb.Entry{}
	for _, r := range rows {
		entries = append(entries, &pb.Entry{
			Id:       r.Entry.ID,
			Day:      r.Entry.Day,
			Title:    r.Entry.Title,
			Comment:  r.Entry.Comment,
			Url:      r.Entry.URL,
			ImageUrl: convertImageURL(r.Entry.ImageURL),
			Owner: &pb.User{
				Id:      r.User.ID,
				Name:    r.User.Name,
				IconUrl: r.User.IconURL,
			},
		})
	}

	return entries, nil
}

func (s *Service) queries() *adventar_db.Queries {
	return s.clients.DB().Queries()
}

func isValidURL(s string) bool {
	u, err := url.Parse(s)
	if err != nil {
		return false
	}

	return u.Scheme == "http" || u.Scheme == "https"
}

func makeSelectValue(data map[string][]string) []string {
	var r []string
	for ns, columns := range data {
		for _, c := range columns {
			v := fmt.Sprintf("%s.%s as `%s.%s`", ns, c, ns, c)
			r = append(r, v)
		}
	}
	return r
}
