package scraper

// ItemNum 获取的数据数量
const ItemNum = 10

// Item 表示爬取的每一条信息
type Item struct {
	Tag     string `json:"tag,omitempty"`
	Title   string `json:"title,omitempty"`
	Link    string `json:"link,omitempty"`
	PubDate string `json:"pub_date,omitempty"`
	IsToday bool   `json:"is_today"`
}

// Scraper 表示每一个具体爬虫应该实现的接口
type Scraper interface {
	Scrape(num int) ([]Item, error)
}
