package scraper

import "testing"

func TestWww(t *testing.T) {
	www := &Www{}
	items, err := www.Scrape(10)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(items)
}
