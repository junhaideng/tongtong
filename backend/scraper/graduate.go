package scraper

import (
	"fmt"
	"github.com/gocolly/colly"
	"strings"
	"time"
)

var graduateLink = "https://yzb.sjtu.edu.cn/index/zkxx/sszs.htm"

// Graduate 爬取研招网通知信息
type Graduate struct {
}

// Scrape 实现通用爬虫接口
func (g *Graduate) Scrape(num int) ([]Item, error) {
	tag := "研究生招生网"
	c := colly.NewCollector()
	items := make([]Item, 0)

	c.OnHTML("div.ny_r > div > div > ul", func(el *colly.HTMLElement) {
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
			date := element.DOM.Find("span:nth-child(2)").Text()
			date = strings.TrimSpace(date)
			parsedDate, err := time.Parse("2006-01-02", date)
			if err != nil {
				fmt.Println(err)
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
	c.Visit(graduateLink)

	return items, nil
}
