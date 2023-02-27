package sql

import (
	"context"
	"goimdemo/common"
	"goimdemo/models"
	"sync"
	"time"
)

// @title:
// @author:
// @param: gorNum int 启动记录的携程数量
// @description: 记录用户发送消息到mysql
// @since:2023-02-24 22:06:36
func RecordToMysqlFunc(ctx context.Context, gorNum int) {
	var wg sync.WaitGroup
	for i := 1; i <= gorNum; i++ {
		wg.Add(1)
		go createMessage(ctx, &wg, i)
	}
	wg.Wait()
}
func createMessage(ctx context.Context, wg *sync.WaitGroup, sleep int) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case msg := <-common.RecordToMysql:
			mMsg := ReplaceMessageData(msg)
			mMsg.SaveMessage()
		default:
			time.Sleep(time.Millisecond * time.Duration(sleep*100))
		}
	}
}
func ReplaceMessageData(m common.Message) *models.Message {
	message := &models.Message{
		FormId:   m.FormId,
		TargetId: m.TargetId,
		GroupId:  m.GroupId,
		Type:     m.Type,
		Media:    m.Media,
		Content:  m.Content, // 消息内容
		Pic:      m.Pic,
		Url:      m.Url,
		Desc:     m.Desc,
		Amount:   m.Amount,
	}
	return message
}
