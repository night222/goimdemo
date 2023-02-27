package utils

import (
	"context"
	"encoding/json"
	"errors"
	"goimdemo/common"
	"strconv"
	"strings"
)

// @Title:message.go
// @Author:proket
// @Description: 对message的操作
// @Since:2023-02-24 23:50:54
func GetGroupByUserId(UserId uint) (string, error) {
	for k, v := range common.GroupByUserId {
		if isBetween(v, "-", UserId) {
			return common.GroupByUserId[k], nil
		}
	}
	return "", errors.New("不在区间内")
}
func GetChanUser(id uint) (chan common.Message, error) {
	var chanUser chan common.Message
	key, err := GetGroupByUserId(id)
	if err != nil {
		return chanUser, err
	}
	//加读锁
	common.MessageQueueChanLock.RLock()
	chanUser, ok := common.MessageQueueChan[key][id]
	if !ok {
		common.MessageQueueChan[key][id] = make(chan common.Message, 10)
		chanUser = common.MessageQueueChan[key][id]
	}
	common.MessageQueueChanLock.RUnlock()
	return chanUser, err
}
func GetGroupByGroupId(GroupId uint) (string, error) {
	for k, v := range common.GroupByGroupId {
		if isBetween(v, "-", GroupId) {
			return common.GroupByGroupId[k], nil
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
		err = GroupWirteMessage(msg)
	case common.PrivateChat:
		err = PrivateWriteMessage(msg, msg.TargetId)
	case common.PublicChat:
		err = PublicWriteMessage(msg)
	}
	return
}

// 群消息
func GroupWirteMessage(m common.Message) (err error) {
	key, err := GetGroupByGroupId(m.GroupId)
	if err != nil {
		return
	}
	cmd := common.RedisClient.HGet(context.Background(), key+common.ConfigData.Group.PostfixGroup, strconv.Itoa(int(m.GroupId)))
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
		err = PrivateWriteMessage(m, uint(targetid))
		if err != nil {
			sLoger.Debugf("failed GroupWriteMessage Groupid:%d TargedId:%d  err:%v \n", m.GroupId, targetid, err)
		}
	}
	return
}

// 私聊消息
func PrivateWriteMessage(m common.Message, TargeId uint) (err error) {
	chanUser, err := GetChanUser(TargeId)
	if err != nil {
		return
	}
	common.RecordToMysql <- m
	chanUser <- m
	return
}

// 广播
func PublicWriteMessage(m common.Message) (err error) {
	sLoger, Loger := GetSugarLogerAndLoger()
	defer Loger.Sync()
	common.MessageQueueChanLock.RLock()
	for _, v := range common.MessageQueueChan {
		for k, v := range v {
			common.RecordToMysql <- m
			v <- m
			if err != nil {
				sLoger.Debugf("failed PublicWriteMessage TargedId:%d  err:%v \n", k, err)
			}
		}
	}
	common.MessageQueueChanLock.RUnlock()
	return
}

// 序列化Message结构体
func EncodeMessageJson(m common.Message) (string, error) {
	data, err := json.Marshal(m)
	return string(data), err
}

// 反序列化
func DecodeMessageJson(data string) (common.Message, error) {
	m := &common.Message{}
	err := json.Unmarshal([]byte(data), m)
	return *m, err
}
