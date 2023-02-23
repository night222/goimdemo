package service

import (
	"goimdemo/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

//token操作

// 校验token
func VaildToken(ctx *gin.Context) {
	u := getTokenUserData(ctx)
	if u.Name == "" {
		ctx.JSON(400, gin.H{
			"message": "Invail Token2",
		})
		ctx.Abort()
	}
}

func getTokenUserData(ctx *gin.Context) utils.TokenClaims {
	token, err := ctx.Cookie("token")
	if err != nil || token == "" {
		//websocker 设置token在请求头 Sec-WebSocket-Protocol
		token = ctx.Request.Header.Get("Sec-Websocket-Protocol")
		token = strings.TrimSpace(token)
		if token == "" {
			return utils.TokenClaims{}
		}
	}
	t, _ := utils.ParseToken(token)
	return t
}
