package initialiaze

import "goimdemo/common"

func InitMessage() {
	common.GroupByUserId = common.ConfigData.Group.UserId
	common.GroupByUserId = common.ConfigData.Group.GroupId
	common.MessageQueueChan = make(map[string]map[uint]chan common.Message, len(common.GroupByUserId))
	for i := 0; i < len(common.GroupByUserId); i++ {
		common.MessageQueueChan[common.GroupByUserId[i]] = make(map[uint]chan common.Message, 1000)
	}
	common.RecordToMysql = make(chan common.Message, 100)
}
