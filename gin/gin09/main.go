package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func indexHanddler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "index",
	})
}
func m1(c *gin.Context) {
	fmt.Println("m1 in ...")
	start := time.Now()
	c.Next() //调用后续的处理函数
	//c.Abort()  组织调用后续处理函数
	cost := time.Since(start)
	fmt.Printf("cost:%v\n", cost)
	fmt.Println("m1 out")
}
func m2(c *gin.Context) {
	fmt.Println("m2 in ...")
	c.Abort() //阻止调用后续处理函数
	return
	fmt.Println("m2 out...")
}
func main() {
	r := gin.Default() //默认使用Logger和Recovery中间件
	//全局注册中间件
	r.Use(m1, m2)
	//r.GET("/index", m1, indexHanddler)
	//r.GET("/shop", m1, func(c *gin.Context)
	r.GET("/index", indexHanddler)
	r.GET("/shop", func(c *gin.Context) {
		fmt.Println("shop")
		c.JSON(http.StatusOK, gin.H{
			"商品": "shop",
		})
	})
	r.Run()
}
