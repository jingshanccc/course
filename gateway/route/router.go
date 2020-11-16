package route

import (
	"course/gateway/handler"
	"course/gateway/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter(service ...interface{}) *gin.Engine {
	ginRouter := gin.Default()
	ginRouter.Use(middleware.InitMiddleware(service))
	v1 := ginRouter.Group("/api/v1")

	v1.GET("/user/list", handler.GetUserList)
	v1.POST("/user/login", handler.Login)
	v1.POST("/user/save-password", handler.SavePassword)
	v1.POST("/user/save", handler.Save)
	return ginRouter
}
