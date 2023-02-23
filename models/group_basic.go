package models

import "gorm.io/gorm"

//@Title:group_basic.go
//@Author:proket
//@Description: 群结构表
//@Since:2023-02-23 20:13:26

type GroupBasic struct {
	gorm.Model
	GroupId   uint   //群id
	GroupName string // 群名称
	Member    string //群成员id
	Icon      string //群头像
	Desc      string
}
