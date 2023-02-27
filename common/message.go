package common

import (
	"sync"
)

// @Title:message.go
// @Author:proket
// @Description: 发送消息的状态等公共设置
// @Since:2023-02-23 20:19:39
type MessageType int
type MessageMedia int

const (
	GroupChat   MessageType = iota //群聊
	PrivateChat                    //私聊
	PublicChat                     //广播
)
const (
	Text  MessageMedia = iota //文字
	Img                       //图片
	Audio                     //音频
)

// @Title:send_message.go
// @Author:proket
// @Description:发送消息包
// @Since:2023-02-24 02:50:13
type Message struct {
	FormId   uint         //发送者
	TargetId uint         //接受者
	GroupId  uint         //接受群
	Type     MessageType  //发送类型
	Media    MessageMedia //消息类型
	Content  string       // 消息内容
	Pic      string
	Url      string
	Desc     string
	Amount   int //其他数字统计
}

// 更具user id 分组 这是接受这id区间
var GroupByUserId = []string{}
var GroupByGroupId = []string{}

// 保存过来的消息
var MessageQueueChanLock sync.RWMutex
var MessageQueueChan map[string]map[uint]chan Message //注意这个是接受者id区间
var RecordToMysql chan Message                        //缓冲保存数据到mysql
