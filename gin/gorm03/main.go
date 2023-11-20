package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Student struct {
	ID   int64
	Name string
	Age  int64
}

func main() {
	//链接数据库
	dsn := "root:88888888@tcp(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	//	把模型与数据库中的表对应起来
	db.AutoMigrate(&Student{})
	//创建
	//st := Student{17, "小王子", sql.NullInt64{0, true}}
	//db.Create(&st)
	//查询
	var st Student
	db.First(&st)
	//fmt.Printf("Student:%#v\n", st)
	//更新
	st.Name = "非凡"
	st.Age = 100
	db.Debug().Save(&st)
	//删除
	var u = Student{}
	u.ID = 10
	u.Name = "非凡"
	db.Debug().Delete(&u)
}
