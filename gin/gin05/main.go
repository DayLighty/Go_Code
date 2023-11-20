package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 创建一个结构体
type UserInFo struct {
	Username string
	Password string
}

func main() {
	//创建路由
	r := gin.Default()
	//加载模版的文件
	r.LoadHTMLGlob("templates/**/*")
	//配置路由
	//获取GET数据展示
	r.GET("/", func(c *gin.Context) {
		id := c.DefaultQuery("id", "1")
		user := c.Query("user")
		pwd := c.Query("pwd")

		c.JSON(http.StatusOK, gin.H{
			"user": user,
			"pwd":  pwd,
			"id":   id,
		})
	})
	//获取post数据展示
	r.GET("/default/user", func(c *gin.Context) {
		c.HTML(http.StatusOK, "default/user.html", gin.H{})
	})
	//获取表单post过来的数据
	r.POST("/doAddUser", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		age := c.DefaultPostForm("age", "20")

		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
			"age":      age,
		})
	})
	//动态路由传值
	r.GET("/list/:cid", func(c *gin.Context) {
		cid := c.Param("cid")
		c.String(200, "%v", cid)
	})
	//获取 GET POST 传递的数据绑定到结构体
	r.GET("/getuser", func(c *gin.Context) {
		//定义结构体变量
		//有个误区
		//user := &UserInFo{}视频教程加了&符号按下面编写没法获取数据
		//user := UserInFo{}
		var user UserInFo
		err := c.ShouldBind(user)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"err": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, user)
		}
		fmt.Println(user)
	})
	//路由分组
	ApiRouters := r.Group("/api")
	{
		ApiRouters.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, "我是一个api接口")
		})
		ApiRouters.GET("/userlist", func(c *gin.Context) {
			c.String(http.StatusOK, "我是一个api接口--userlist")
		})
		ApiRouters.GET("/plist", func(c *gin.Context) {
			c.String(http.StatusOK, "我是一个api接口--plist")
		})
	}
	r.Run()
}
