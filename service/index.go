package service

import "github.com/gin-gonic/gin"

// GetIndex
// @Tags 首页
// @Success 200 {string}  welcome
// @Router /index [get]
func Index(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"massage": "welcome"})
}
