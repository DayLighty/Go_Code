package service

import (
	"fmt"
	"ginchat/models"
	"ginchat/utils"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	_ "golang.org/x/net/websocket"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

// GetUserList
// @Summary 所有用户
// @Tags 用户模块
// @Success 200 {string} json{"code","message"}
// @Router /user/getUserList [get]
func GetUserList(c *gin.Context) {
	data := make([]*models.UserBasic, 10)
	data = models.GerUserList()
	c.JSON(200, gin.H{
		"message": data,
	})
}

// CreateUser
// @Summary 新增用户
// @Tags 用户模块
// @param name query string false "用户名"
// @param password query string false "密码"
// @param repassword query string false "确认密码"
// @Success 200 {string} json{"code","message"}
// @Router /user/createUser [get]
func CreateUser(c *gin.Context) {
	user := models.UserBasic{}
	user.Name = c.Query("name")
	password := c.Query("password")
	repassword := c.Query("repassword")
	salt := fmt.Sprintf("%06d", rand.Int31())
	user.Salt = salt
	data := models.FindUserByName(user.Name)
	if user.Name == "" {
		c.JSON(200, gin.H{
			"code":    -1, //0 成功  -1 失败
			"message": "用户名为空，请输入用户名!",
		})
		return
	}
	if password == "" {
		c.JSON(200, gin.H{
			"code":    -1, //0 成功  -1 失败
			"message": "密码为空，请输入密码!",
		})
		return
	}
	if repassword == "" {
		c.JSON(200, gin.H{
			"code":    -1, //0 成功  -1 失败
			"message": "确认密码为空，请输入密码!",
		})
		return
	}
	if data.Name != "" {
		c.JSON(200, gin.H{
			"code":    -1, //0 成功  -1 失败
			"message": "用户名已注册",
		})
		return
	}
	if password != repassword {
		c.JSON(200, gin.H{
			"code":    -1, //0 成功  -1 失败
			"message": "两次密码不一致",
		})
		return
	}
	//user.PassWord = password
	user.PassWord = utils.MakePassword(password, salt)
	models.CreateUser(user)
	c.JSON(200, gin.H{
		"code":    0, //0 成功  -1 失败
		"message": "用户新增成功",
		"data":    user,
	})
}

// FindUserByNameAndPwd
// @Summary 登录用户
// @Tags 用户模块
// @param name query string false "用户名"
// @param password query string false "密码"
// @Success 200 {string} json{"code","message"}
// @Router /user/findUserByNameAndPwd [post]
func FindUserByNameAndPwd(c *gin.Context) {
	data := models.UserBasic{}
	//name := c.Query("name")
	//password := c.Query("password")
	name := c.Request.FormValue("name")
	password := c.Request.FormValue("password")
	fmt.Println(name, password)
	user := models.FindUserByName(name)
	if user.Name == "" {
		c.JSON(200, gin.H{
			"code":    -1, //0 成功  -1 失败
			"message": "该用户不存在",
			"data":    data,
		})
	}

	flag := utils.ValidPassword(password, user.Salt, user.PassWord)
	if !flag {
		c.JSON(200, gin.H{
			"code":    -1, //0 成功  -1 失败
			"message": "密码不正确",
			"data":    data,
		})
	}
	pwd := utils.MakePassword(password, user.Salt)
	data = models.FindUserByNameAndPwd(name, pwd)
	c.JSON(200, gin.H{
		"code":    0, //0 成功  -1 失败
		"message": "登陆成功",
		"data":    data,
	})
}

// DeleteUser
// @Summary 删除用户
// @Tags 用户模块
// @param id query string false "id"
// @Success 200 {string} json{"code","message"}
// @Router /user/deleteUser [get]
func DeleteUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.Query("id"))
	user.ID = uint(id)
	models.DeleteUser(user)
	c.JSON(200, gin.H{
		"code":    0, //0 成功  -1 失败
		"message": "用户删除成功",
		"data":    user,
	})
}

// UpdateUser
// @Summary 修改用户
// @Tags 用户模块
// @Param id formData string false "id"
// @Param name formData string false "name"
// @Param password formData string false "password"
// @Param phone formData string false "phone"
// @Param email formData string false "email"
// @Success 200 {string} json{"code","message"}
// @Router /user/updateUser [post]
func UpdateUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.PostForm("id"))
	user.ID = uint(id)
	user.Name = c.PostForm("name")
	user.PassWord = c.PostForm("password")
	user.Phone = c.PostForm("phone")
	user.Email = c.PostForm("email")
	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{
			"code":    -1, //0 成功  -1 失败
			"message": "用户修改失败",
			"data":    user,
		})
	} else {
		models.UpdateUser(user)
		c.JSON(200, gin.H{
			"code":    0, //0 成功  -1 失败
			"message": "用户修改成功",
			"data":    user,
		})
	}
}

// 防止跨域站点伪造请求
var upGrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// 发送消息准备工作
func SendMsg(c *gin.Context) {
	ws, err := upGrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(ws *websocket.Conn) {
		err = ws.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(ws)
	MsgHandler(c, ws)
}

func MsgHandler(c *gin.Context, ws *websocket.Conn) {
	msg, err := utils.Subscribe(c, utils.PublishKey)
	if err != nil {
		fmt.Println("发送消息失败", err)
		return
	}
	fmt.Println("发送消息...", msg)

	tm := time.Now().Format("2006-01-02 15:04:05")
	m := fmt.Sprintf("[ws][%s]:%s", tm, msg)
	err = ws.WriteMessage(1, []byte(m))
	if err != nil {
		fmt.Println(err)
	}
}

func SenUserdMsg(c *gin.Context) {
	models.Chat(c.Writer, c.Request)
}
