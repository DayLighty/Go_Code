package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	//Any: 请求方法的大集合
	r.Any("/user", func(c *gin.Context) {
		switch c.Request.Method {
		case http.MethodGet:
			c.JSON(http.StatusOK, gin.H{"message": "Get"})
		case http.MethodPost:
			c.JSON(http.StatusOK, gin.H{"message": "Post"})
		case http.MethodPut:
			c.JSON(http.StatusOK, gin.H{"message": "Put"})
		case http.MethodDelete:
			c.JSON(http.StatusOK, gin.H{"message": "Delete"})
		}
	})
	//NoRoute 没有定义的路由
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"msg": "此路由没有定义"})
	})
	//路由组
	//把公用的前缀提取出来，创建一个路由组
	videoGroup := r.Group("/video")
	{
		videoGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "/video/index"})
		})
		videoGroup.GET("/xx", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "/video/xx"})
		})
		videoGroup.GET("/oo", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "/video/oo"})
		})
	}
	//视频的首页和详情页
	//r.GET("/video/index", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{"msg": "/video/index"})
	//})
	//r.GET("/video/xx", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{"msg": "/video/xx"})
	//})
	//r.GET("/video/oo", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{"msg": "/video/oo"})
	//})
	r.Run()

}
