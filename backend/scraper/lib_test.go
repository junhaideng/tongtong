package scraper

import "testing"

func TestLib(t *testing.T) {
	lib := &Lib{}
	items, err := lib.Scrape(10)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(items)
}
