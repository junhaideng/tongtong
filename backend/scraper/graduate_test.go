package scraper

import "testing"

func TestGraduate(t *testing.T) {
	g := &Graduate{}
	items, err := g.Scrape(10)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v\n", items)
}
