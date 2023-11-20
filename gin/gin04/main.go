package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"time"
)

// 时间转换成日期
func UnixToTime(timestamp int) string {
	fmt.Println(timestamp)
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05") //格式化
}
func Println(str1 string, str2 string) string {
	fmt.Println(str1, str2)
	return str1 + "---" + str2
}
func main() {
	//创建路由
	r := gin.Default()
	//自定义模版函数  注意把这个函数要放在加载模版前
	//template.FuncMap
	r.SetFuncMap(template.FuncMap{
		"UnixToTime": UnixToTime, //注册模版函数,,使用模版函数必须注册，不然找不到会报错
		"Println":    Println,
	})
	//加载模版的文件
	r.LoadHTMLGlob("template/**/*")
	//配置路由
	//配置静态web目录，第一个参数表示路由，第二个参数表示映射目录
	r.Static("/static", "./static")
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
			"name":  "小王",
			"price": 20,
			"date":  1697808712,
		})
	})

	r.Run()
}
