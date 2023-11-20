package main

import (
	"ginchat/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	//连接数据库
	dsn := "root:88888888@tcp(127.0.0.1:3306)/ginchat?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	//	迁移 schema
	db.AutoMigrate(&models.UserBasic{})
	db.AutoMigrate(&models.Message{})
	db.AutoMigrate(&models.Contact{})
	db.AutoMigrate(&models.Gropbasic{})

	////	Create
	//user := &models.UserBasic{}
	//user.Name = "马伯林"
	//db.Create(user)
	////	read
	//fmt.Println(db.First(user))
	////	update
	//db.Model(user).Update("PassWord", "66666")
	//models.GerUserList()

}
