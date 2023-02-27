package goimmessage

import "encoding/json"

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

// 序列化Message结构体
func EncodeMessageJson(m Message) (string, error) {
	data, err := json.Marshal(m)
	return string(data), err
}

// 反序列化
func DecodeMessageJson(data string) (Message, error) {
	m := &Message{}
	err := json.Unmarshal([]byte(data), m)
	return *m, err
}
