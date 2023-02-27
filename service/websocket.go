package service

import (
	"context"
	"goimdemo/common"
	"goimdemo/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// websocket操作
var upgrader = websocket.Upgrader{
	//解决跨域问题
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// SendWeb
// PingExample godoc
// @Tags websocket 测试接口
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /sendweb [websocket]
func SendWeb(ctx *gin.Context) {
	sLoger, loger := utils.GetSugarLogerAndLoger()
	defer loger.Sync()
	token := getTokenUserData(ctx)
	if token.Name == "" {
		sLoger.Debugln("Invail Token")
		return
	}
	ws, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	//关闭连接
	defer ws.Close()
	if err != nil {
		sLoger.Debugf("failed upgrader.Upgrade ip:%s url:%s",
			ctx.ClientIP(),
			ctx.Request.URL.Host,
		)
		return
	}
	context, cancel := context.WithCancel(context.Background())
	chanErr := make(chan error, 1)
	go WriteMessage(ctx, chanErr, token.ID, ws)
	for {
		if len(chanErr) == 1 {
			cancel()
			sLoger.Debugln("write:", err)
			break
		}
		_, msg, err := ws.ReadMessage()
		if err != nil {
			cancel()
			sLoger.Debugln("read:", err)
			break
		}
		//发布msg到redis
		key, err := utils.GetGroupByUserId(token.ID)
		if err != nil {
			cancel()
			sLoger.Debugln("websocker common.GetGroupByUserId", err)
			return
		}
		err = utils.Publist(context, key+common.ConfigData.Group.PostfixUser, string(msg))
		if err != nil {
			cancel()
			sLoger.Debugln("failed redis publist sendweb", err)
			break
		}
	}
}

// 接收消息
func WriteMessage(ctx context.Context, chanErr chan error, id uint, ws *websocket.Conn) {
	chanMsg, err := utils.GetChanUser(id)
	if err != nil {
		chanErr <- err
		return
	}
	for {
		select {
		case <-ctx.Done():
			return
		case user := <-chanMsg:
			err = ws.WriteJSON(user)
			if err != nil {
				chanErr <- err
				return
			}
		}
	}
}
