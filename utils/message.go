package utils

import (
	"context"
	"encoding/json"
	"errors"
	"goimdemo/common"
	"strconv"
	"strings"
	"sync"
)

// @Title:send_message.go
// @Author:proket
// @Description:发送消息包
// @Since:2023-02-24 02:50:13
type Message struct {
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

// 更具user id 分组 这是接受这id区间
var GroupByUserId = []string{}
var GroupByGroupId = []string{}

// 保存过来的消息
var MessageQueueChanLock sync.RWMutex
var MessageQueueChan map[string]map[uint]chan Message //注意这个是接受者id区间
func GetGroupByUserId(UserId uint) (string, error) {
	for k, v := range GroupByUserId {
		if isBetween(v, "-", UserId) {
			return GroupByUserId[k], nil
		}
	}
	return "", errors.New("不在区间内")
}
func GetChanUser(id uint) (chan Message, error) {
	var chanUser chan Message
	key, err := GetGroupByUserId(id)
	if err != nil {
		return chanUser, err
	}
	//加读锁
	MessageQueueChanLock.RLock()
	chanUser, ok := MessageQueueChan[key][id]
	if !ok {
		MessageQueueChan[key][id] = make(chan Message, 10)
		chanUser = MessageQueueChan[key][id]
	}
	MessageQueueChanLock.RUnlock()
	return chanUser, err
}
func GetGroupByGroupId(GroupId uint) (string, error) {
	for k, v := range GroupByGroupId {
		if isBetween(v, "-", GroupId) {
			return GroupByGroupId[k], nil
		}
	}
	return "", errors.New("群不在区间内")
}

// @title:
// @author:
// @param: name type description
// @return: name type description
// @description: 获取用户对应消息接收的管道
// @since:2023-02-23 21:42:05
func isBetween(between, sep string, num uint) bool {
	index := strings.Index(between, sep)
	if index == -1 {
		return false
	}
	minS, maxS := between[:index], between[index+1:]
	max, err := strconv.ParseUint(maxS, 10, 0)
	if err != nil {
		return false
	}
	min, err := strconv.ParseUint(minS, 10, 0)
	if err != nil {
		return false
	}
	if num <= uint(max) && num >= uint(min) {
		return true
	}
	return false
}

// 推送消息
func WriteMessage(data string) (err error) {
	msg, err := DecodeMessageJson(data)
	if err != nil {
		return
	}
	switch msg.Type {
	case common.GroupChat:
		err = msg.GroupWirteMessage()
	case common.PrivateChat:
		err = msg.PrivateWriteMessage(msg.TargetId)
	case common.PublicChat:
		err = msg.PublicWriteMessage()
	}
	return
}

// 群消息
func (m *Message) GroupWirteMessage() (err error) {
	key, err := GetGroupByGroupId(m.GroupId)
	if err != nil {
		return
	}
	cmd := RedisClient.HGet(context.Background(), key+ConfigData.Group.PostfixGroup, strconv.Itoa(int(m.GroupId)))
	data, err := cmd.Result()
	if err != nil {
		return
	}
	sLoger, Loger := GetSugarLogerAndLoger()
	defer Loger.Sync()
	GroupIds := strings.Split(data, ",")
	for _, v := range GroupIds {
		targetid, _ := strconv.ParseUint(v, 10, 0)
		if m.FormId == uint(targetid) {
			continue
		}
		err = m.PrivateWriteMessage(uint(targetid))
		if err != nil {
			sLoger.Debugf("failed GroupWriteMessage Groupid:%d TargedId:%d  err:%v \n", m.GroupId, targetid, err)
		}
	}
	return
}

// 私聊消息
func (m *Message) PrivateWriteMessage(TargeId uint) (err error) {
	chanUser, err := GetChanUser(TargeId)
	if err != nil {
		return
	}
	chanUser <- *m
	return
}

// 广播
func (m *Message) PublicWriteMessage() (err error) {
	sLoger, Loger := GetSugarLogerAndLoger()
	defer Loger.Sync()
	MessageQueueChanLock.RLock()
	for _, v := range MessageQueueChan {
		for k, v := range v {
			v <- *m
			if err != nil {
				sLoger.Debugf("failed PublicWriteMessage TargedId:%d  err:%v \n", k, err)
			}
		}
	}
	MessageQueueChanLock.RUnlock()
	return
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
