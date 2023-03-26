package util_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/adventar/adventar/backend/pkg/util"
)

func TestSiteTitle(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "<head><title>foo</title><title>bar</title><body><svg><title>baz</title></svg></body>")
	})
	ts := httptest.NewServer(h)
	defer ts.Close()

	fetcher := util.SiteMetaFetcher{}
	res, err := fetcher.Fetch(ts.URL)
	if err != nil {
		t.Fatal(err)
	}

	if res.Title != "foo" {
		t.Errorf("actual: %s, expected: foo", res.Title)
	}
}
