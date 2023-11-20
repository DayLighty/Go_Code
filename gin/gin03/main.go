package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//创建路由
	r := gin.Default()
	//配置模版的文件
	r.LoadHTMLGlob("templates/**/*")
	//配置路由
	//后台
	r.GET("/default/news", func(c *gin.Context) {
		c.HTML(http.StatusOK, "default/news.html", gin.H{
			"title": "我是后台新闻",
		})
	})
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "default/index.html", gin.H{
			"title": "我是商品页面后台",
			"price": 20,
		})
	})
	//前台
	r.GET("/admin/news", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin/news.html", gin.H{
			"title": "我是前台新闻",
		})
	})
	r.GET("admin", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin/index.html", gin.H{
			"title": "我是商品页面前台",
			"price": 20,
		})
	})

	r.Run()
}
