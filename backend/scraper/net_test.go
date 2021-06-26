package scraper

import "testing"

func TestNet(t *testing.T) {
	n := &Net{}
	items, err := n.Scrape(ItemNum)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(items)

}
