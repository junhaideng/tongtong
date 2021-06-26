package scraper

import (
	"github.com/gocolly/colly"
	"time"
)

var wwwLink = "https://www.sjtu.edu.cn/tg/index.html"

// Www 爬取交大主页中的内容
type Www struct {
}

// Scrape 实现通用爬虫接口
func (w *Www) Scrape(num int) ([]Item, error) {
	c := colly.NewCollector()
	items := make([]Item, 0)
	tag := "交大主页"
	c.OnHTML("#main > section.pageMain > div > div > ul", func(el *colly.HTMLElement) {
		el.ForEachWithBreak("li", func(i int, element *colly.HTMLElement) bool {
			// 如果达到了需要的数量，那么停止爬取
			if len(items) == num {
				return false
			}
			// 通知标题
			title := element.DOM.Find("a").Text()
			// 通知信息的详情链接
			link := el.Request.AbsoluteURL(element.DOM.Find("a").AttrOr("href", ""))
			// 通知发布的日期
			date := element.DOM.Find("span").Text()

			parsedDate, err := time.Parse("2006.01.02", date)
			if err != nil {
				return true
			}
			// 添加到结果集
			items = append(items, Item{
				Tag:     tag,
				Title:   title,
				Link:    link,
				PubDate: formatTime(parsedDate),
				IsToday: isToday(parsedDate),
			})
			return true
		})
	})

	// 访问网页，爬取数据
	c.Visit(wwwLink)

	return items, nil
}
