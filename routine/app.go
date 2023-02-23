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
	r.POST("/login", service.Login)
	r.GET("/sendweb", service.SendWeb)
	user := r.Group("/user")
	user.Use(service.VaildToken)
	{
		user.GET("", service.UserList)
		user.POST("", service.CreateUser)
		user.DELETE("", service.DeleteUser)
		user.PUT("", service.UpdateUser)
	}
	return r
}
