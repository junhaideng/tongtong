package scraper

import (
	"testing"
)

func TestJwc(t *testing.T) {
	jwc := &Jwc{}
	items, err := jwc.Scrape(ItemNum)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(items)
}
