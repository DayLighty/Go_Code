package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GET 请求
func getkey(c *gin.Context) {
	s := c.Query("name")
	age := c.DefaultQuery("age", "xxoo")
	c.String(http.StatusOK, "this is getkey()获取到key=%s  %s", s, age)
}

// POST 请求
func posrVal(c *gin.Context) {
	//s := c.PostForm("name")
	//age := c.DefaultPostForm("age", "xxoo")
	c.String(http.StatusOK, "我是post请求")
}
func main() {
	//创建一个默认路由
	r := gin.Default()
	//配置路由

	//gin.H 与 map[string]interface 本质一样
	r.GET("ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{ //返回JSON格式数据   type H map[string]any
			"msg": "pong",
		})
	})
	r.GET("/news", func(c *gin.Context) {
		c.String(http.StatusOK, "这是我们的新闻页面")
	})
	r.POST("/postVal", posrVal)
	//r.GET("/someGet", someGet)
	r.GET("/getkey", getkey)
	r.Run()
}
