package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		//c.JSON(http.StatusOK, gin.H{
		//	"status": "ok",
		//})
		//重定向  c.Redirect
		//跳转到 指定URL
		c.Redirect(http.StatusMovedPermanently, "https://github.com/coders-daybreak")
	})
	//路由重定向
	r.GET("/a", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "a1",
		})
		//跳转到 /b 对应的路由处理函数
		c.Request.URL.Path = "/b" //把请求的URL修改
		r.HandleContext(c)        //继续后续的结果
		c.JSON(http.StatusOK, gin.H{
			"message": "a2",
		})
	})
	r.GET("/b", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "b",
		})
	})
	r.Run()
}
