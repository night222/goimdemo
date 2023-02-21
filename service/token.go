package service

import (
	"goimdemo/utils"

	"github.com/gin-gonic/gin"
)

//token操作

// 校验token
func VaildToken(ctx *gin.Context) {
	token, err := ctx.Cookie("token")
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "Invail Token",
		})
		ctx.Abort()
		return
	}
	t, _ := utils.ParseToken(token)
	if t.Name == "" {
		ctx.JSON(400, gin.H{
			"message": "Invail Token2",
		})
		ctx.Abort()
	}
}
