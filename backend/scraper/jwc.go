package scraper

import (
	"fmt"
	"github.com/gocolly/colly"
	"strings"
	"time"
)

var jwcLink = "https://jwc.sjtu.edu.cn/xwtg/tztg.htm"

// Jwc 用于获取教务处信息
type Jwc struct {
}

// Scrape 实现爬虫通用接口
func (j *Jwc) Scrape(num int) ([]Item, error) {
	tag := "本科教学信息网"
	// 创建一个collector
	c := colly.NewCollector()
	// 用于保存数据
	items := make([]Item, 0)
	// 界面中所有的通知信息所在的selector
	queryString := "div.ny_right_con > div > ul"
	// 注册一个钩子
	c.OnHTML(queryString, func(element *colly.HTMLElement) {
		element.ForEachWithBreak("li", func(i int, element *colly.HTMLElement) bool {
			// 仅获取需要的数量，如果需要的过多，仅会返回当前界面仅有的
			if len(items) == num {
				return false
			}
			link := element.DOM.Find("a")
			// 发表日期
			day := strings.TrimSpace(element.DOM.Find("div.sj > h2").Text())
			date := strings.TrimSpace(element.DOM.Find("div.sj > p").Text()) + "."+day
			// 去掉两边的[]，仅获取到日期字符串
			parsedDate, err := time.Parse("2006.01.02", date)
			if err != nil {
				fmt.Println(err)
				return true
			}
			// 添加一条新的数据到items中
			items = append(items, Item{
				Tag:     tag,
				Title:   strings.TrimSpace(link.Text()),
				Link:    element.Request.AbsoluteURL(link.AttrOr("href", "")),
				PubDate: date,
				IsToday: isToday(parsedDate),
			})
			return true
		})
	})
	// 开始访问该网站
	c.Visit(jwcLink)
	return items, nil
}

