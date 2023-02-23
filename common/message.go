package common

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
