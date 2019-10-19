package util

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type SiteMetaFetcher struct{}

// SiteMeta represents site meta infomation.
type SiteMeta struct {
	Title    string
	ImageURL string
}

// TODO: Support charset other than utf8
func (smf *SiteMetaFetcher) Fetch(url string) (*SiteMeta, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	meta := &SiteMeta{}
	doc.Find("title").Each(func(i int, s *goquery.Selection) {
		meta.Title = s.Text()
	})
	doc.Find("meta[name='twitter:image'], meta[property='og:image']").Each(func(i int, s *goquery.Selection) {
		c := s.AttrOr("content", "")
		if c != "" {
			meta.ImageURL = c
		}
	})

	return meta, err
}
