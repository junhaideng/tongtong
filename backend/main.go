package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sjtu-notice/middleware"
	"sjtu-notice/scraper"
)

// ReqQuery http请求参数
type ReqQuery struct {
	// 请求的通知信息数量
	Num int `form:"limit"`
}

func main() {
	r := gin.Default()
	// 跨域请求
	r.Use(middleware.Cors())
	// 健康检查接口
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": SuccessCode,
			"msg":  Healthy,
		})
	})

	// 设置一个路由组
	api := r.Group("/api")

	// 路由
	api.GET("/notice/:from", func(c *gin.Context) {
		// 通知来源渠道
		from := c.Param("from")

		// 绑定请求参数，绑定失败，返回对应的相应提示
		query := ReqQuery{}
		if err := c.ShouldBind(&query); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": ErrCode,
				"msg":  ErrBadRequest.Error(),
				"data": []int{},
			})
			return
		}
		// 如果没有设置请求的通知数量，设置为默认的
		if query.Num == 0 {
			query.Num = DefaultQueryNum
		}

		// 通知信息
		var items []scraper.Item
		var err error

		// 判断获取哪个通知信息渠道来源
		switch from {
		case "jwc":
			jwc := &scraper.Jwc{}
			items, err = jwc.Scrape(query.Num)
		case "net":
			net := &scraper.Net{}
			items, err = net.Scrape(query.Num)
		case "lib":
			lib := &scraper.Lib{}
			items, err = lib.Scrape(query.Num)
		case "www":
			www := &scraper.Www{}
			items, err = www.Scrape(query.Num)
		case "graduate":
			graduate := &scraper.Graduate{}
			items, err = graduate.Scrape(query.Num)

		// 如果都不能进行匹配，那么不存在该渠道
		default:
			c.JSON(http.StatusBadRequest, gin.H{
				"code": ErrCode,
				"msg":  ErrBadRequest.Error(),
				"data": []int{},
			})
			return
		}

		// 爬取信息的时候出现错误
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": ErrCode,
				"msg":  ErrFailedGet.Error(),
				"data": []int{},
			})
			return
		}

		// 爬取成功，返回对应的内容
		c.JSON(http.StatusOK, gin.H{
			"code": SuccessCode,
			"msg":  Success,
			"data": items,
		})
	})

	// 运行后端服务器
	r.Run()
}
