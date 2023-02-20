package routine

import (
	"goimdemo/docs"
	"goimdemo/service"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = ""
	//swagger 中间件使用
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.GET("/index", service.Index)
	r.GET("/user", service.UserList)
	r.POST("/user", service.CreateUser)
	r.DELETE("/user", service.DeleteUser)
	r.PUT("/user", service.UpdateUser)
	return r
}
