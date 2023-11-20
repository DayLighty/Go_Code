package models

import "gorm.io/gorm"

// 人员关系表
type Contact struct {
	gorm.Model
	OwnerId  uint //谁的关系
	TargetId uint //对应的谁
	Type     int  //对应的类型0 1
	Desc     string
}

func (table *Contact) TableName() string {

	return "contact"
}
