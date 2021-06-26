package scraper

import (
	"github.com/gocolly/colly"
	"time"
)

var libLink = "http://www.lib.sjtu.edu.cn/f/content/list.shtml?Lid=3&lang=zh-cn"

// Lib 爬取图书馆通知信息
type Lib struct {
}

// Scrape 实现通用爬虫接口
func (l *Lib) Scrape(num int) ([]Item, error) {
	c := colly.NewCollector()
	items := make([]Item, 0)
	tag := "图书馆"

	c.OnHTML("div.result_div_list > div > ul", func(el *colly.HTMLElement) {
		el.ForEachWithBreak("li>div > table > tbody", func(i int, element *colly.HTMLElement) bool {
			if element.DOM.Find("tr:nth-child(2)").Nodes == nil {
				return true
			}
			// 如果达到了需要的数量，那么停止爬取
			if len(items) == num {
				return false
			}
			// 通知标题
			title := element.DOM.Find("div.resource_content_title").Text()
			// 通知发布的日期
			date := element.DOM.Find("div.resource_content_time").Text()
			// 通知信息的详情链接
			link := element.DOM.Find("div.resource_content_more>a").AttrOr("href", "")
			link = element.Request.AbsoluteURL(link)

			parsedDate, err := time.Parse("2006-01-02", date)
			if err != nil {
				return true
			}
			// 添加到结果集
			items = append(items, Item{
				Tag:     tag,
				Title:   title,
				Link:    link,
				PubDate: date,
				IsToday: isToday(parsedDate),
			})
			return true
		})
	})
	// 访问网页，爬取数据
	c.Visit(libLink)
	return items, nil
}
