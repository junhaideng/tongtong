package scraper

import (
	"github.com/gocolly/colly"
	"time"
)

var netLink = "https://net.sjtu.edu.cn/xwgg/zxgg.htm"

// Net 爬取网络信息中心通知
type Net struct {
}

// Scrape 实现爬虫接口
func (n *Net) Scrape(num int) ([]Item, error) {
	tag := "网络信息中心"
	items := make([]Item, 0)
	c := colly.NewCollector()
	c.OnHTML("div.main-right > div.right-nr > ul", func(e *colly.HTMLElement) {
		e.ForEachWithBreak("li", func(index int, li *colly.HTMLElement) bool {
			if len(items) == num {
				return false
			}
			// a 标签
			a := li.DOM.Find("a")
			// span 标签
			span := li.DOM.Find("span")
			parsedDate, err := time.Parse("2006-01-02", span.Text())
			if err != nil {
				return true
			}
			// 添加到结果集
			items = append(items, Item{
				Tag:     tag,
				Link:    e.Request.AbsoluteURL(a.AttrOr("href", "")),
				Title:   a.Text(),
				PubDate: span.Text(),
				IsToday: isToday(parsedDate),
			})
			return true
		})
	})
	// 访问网页，爬取数据
	c.Visit(netLink)

	return items, nil
}
