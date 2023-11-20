package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type UserInfo struct {
	ID     uint
	Name   string
	Gender string
	Hobby  string
}

func main() {
	dsn := "root:88888888@tcp(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// 创建表 自动迁移 把结构体和数据表进行对应
	db.AutoMigrate(&UserInfo{})
	//创建数据行
	u1 := UserInfo{1, "daylighty", "男", "篮球"}
	db.Create(&u1)
	//查询
	//var u UserInfo
	//db.First(&u) //查询表中第一条数据保存到u中
	//fmt.Println(u)
	//	更新
	//db.Model(&u).Update("hobby", "双色球")
	////	删除
	//db.Delete(&u)
}
