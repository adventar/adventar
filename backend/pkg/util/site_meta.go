package util

import (
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/m-mizutani/goerr"
)

type SiteMetaFetcher struct{}

// SiteMeta represents site meta infomation.
type SiteMeta struct {
	Title    string
	ImageURL string
}

// TODO: Support charset other than utf8
func (smf *SiteMetaFetcher) Fetch(url string) (*SiteMeta, error) {
	client := http.Client{Timeout: 2 * time.Second}
	res, err := client.Get(url)
	if err != nil {
		return nil, goerr.Wrap(err, "Failed to fetch site meta")
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, goerr.Wrap(err, "Failed to fetch site meta request").
			With("status_code", res.StatusCode).
			With("status", res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, goerr.Wrap(err, "Failed to parse html")
	}

	meta := &SiteMeta{}
	meta.Title = doc.Find("head title").First().Text()
	doc.Find("meta[name='twitter:image'], meta[property='og:image']").Each(func(i int, s *goquery.Selection) {
		c := s.AttrOr("content", "")
		if c != "" {
			meta.ImageURL = c
		}
	})

	return meta, nil
}
