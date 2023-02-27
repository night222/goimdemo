package utils

import (
	"context"
	"goimdemo/common"
	"sync"
)

//发布消息

func Publist(ctx context.Context, channel string, msg string) (err error) {
	i := common.RedisClient.Publish(ctx, channel, msg)
	err = i.Err()
	return
}

// 启动协程订阅消息
// 订阅消息
func Subscribes(ctx context.Context, num int) {
	var wg sync.WaitGroup
	for _, v := range common.GroupByUserId {
		for i := 1; i < num; i++ {
			wg.Add(1)
			Subscribe(ctx, v+common.ConfigData.Group.PostfixUser, &wg)
		}
	}
	wg.Wait()
}
func Subscribe(ctx context.Context, channel string, wg *sync.WaitGroup) {
	defer wg.Done()
	sLoger, Loger := GetSugarLogerAndLoger()
	defer Loger.Sync()
	pub_shub := common.RedisClient.Subscribe(ctx, channel)
	err_num := 0
	for {
		select {
		case <-ctx.Done():
			return
		default:
			msg, err := pub_shub.ReceiveMessage(ctx)
			if err != nil {
				sLoger.Debugln("Subscribe err", err)
				if err_num >= 10 {
					return
				}
				err_num++
			}
			err = WriteMessage(msg.Payload)
			if err != nil {
				sLoger.Debugln("common.WriteMessage(msg.Payload) err", err)
				if err_num >= 10 {
					return
				}
				err_num++
			}
		}
	}
}
