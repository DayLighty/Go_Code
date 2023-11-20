package models

import "gorm.io/gorm"

// 群信息
type Gropbasic struct {
	gorm.Model
	Name    string
	OwnerId uint
	Icon    string
	Desc    string
	Type    int
}

func (table *Gropbasic) TableName() string {

	return "Grop_basic"
}
