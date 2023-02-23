package models

//@Title:message.go
//@Author:proket
//@Description: 消息结构
//@Since:2023-02-23 20:11:58
import (
	"goimdemo/common"
	"goimdemo/utils"

	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	FormId   uint                //发送者
	TargetId uint                //接受者
	GroupId  uint                //接受群
	Type     common.MessageType  //发送类型
	Media    common.MessageMedia //消息类型
	Content  string              // 消息内容
	Pic      string
	Url      string
	Desc     string
	Amount   int //其他数字统计
}

// 定义表名
func (table *Message) TableName() string {
	return "message"
}

// 添加
func SaveMessage(m *Message) *gorm.DB {
	db := utils.DB.Create(m)
	return db
}
