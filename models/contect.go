package models

import (
	"goimdemo/common"

	"gorm.io/gorm"
)

// @Title:contect.go
// @Author:proket
// @Description: 人员关系
// @Since:2023-02-23 20:11:41
type Contect struct {
	gorm.Model
	OwerId  uint               //谁的关系
	TargeId uint               // 对应的谁
	Type    common.ContectType //什么关系
	Desc    string
}

func (c *Contect) TableName() string {
	return "contect"
}
