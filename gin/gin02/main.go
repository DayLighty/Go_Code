package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Students struct {
	Name  string
	Age   int
	Scoer int
}

func main() {
	//创建路由
	r := gin.Default()
	//配置模版的文件
	//r.LoadHTMLFiles("templates/goods.html")
	r.LoadHTMLGlob("templates/*")
	//配置路由
	r.GET("/", func(c *gin.Context) {
		stu := &Students{
			Name:  "小王",
			Age:   18,
			Scoer: 100,
		}
		c.JSON(http.StatusOK, stu)
	})
	//注意：使用c.HTML渲染模版前需要使用LoadHTMLGlob()或者LoadHTMLFiles()方法加载模版
	r.GET("/news", func(c *gin.Context) {
		c.HTML(http.StatusOK, "news.html", gin.H{
			"title": "我是后台数据",
		})
	})
	r.GET("/goods", func(c *gin.Context) {
		c.HTML(http.StatusOK, "goods.html", gin.H{
			"title": "我是商品页面后台",
			"price": 20,
		})
	})
	r.Run()
}
