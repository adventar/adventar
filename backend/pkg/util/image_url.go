package util

import (
	"crypto/sha1"
	"fmt"
	"net/url"
	"os"
)

func ResizableImageURL(imageURL string) string {
	endpoint := os.Getenv("IMAGE_SERVER_ENDPOINT")
	if endpoint == "" || imageURL == "" {
		return imageURL
	}
	salt := os.Getenv("IMAGE_DIGEST_SALT")
	h := sha1.New()
	h.Write([]byte(imageURL + salt))

	return fmt.Sprintf("%s/%x?url=%s", endpoint, h.Sum(nil), url.QueryEscape(imageURL))
}
