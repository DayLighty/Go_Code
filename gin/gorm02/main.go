package main

import (
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Name         string
	Age          sql.NullInt64 //零值类型
	Birthday     *time.Time
	Email        string  `gorm:"type:varchar(100);unique_index"`
	Role         string  `gorm:"size:255"`        // 设置字段大小为255
	MemberNumber *string `gorm:"unique;not null"` // 设置会员号（member number）唯一并且不为空
	Num          int     `gorm:"AUTO_INCREMENT"`  // 设置 num 为自增类型
	Address      string  `gorm:"index:addr"`      // 给address字段创建名为addr的索引
	IgnoreMe     int     `gorm:"-"`               // 忽略本字段
}
type Animal struct {
	AnimalID int64 `gorm:"primary_key"`
	Name     string
	Age      int
}

// 设置Animal   重新建一个daylinghty表内容与Animal一样
func (Animal) TableName() string {
	return "daylinghty"
}

func main() {
	//链接MySQL数据库
	dsn := "root:88888888@tcp(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	//自动创建表
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Animal{})
}
