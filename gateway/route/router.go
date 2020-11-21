package route

import (
	"course/gateway/handler"
	"course/public"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter(service ...interface{}) *gin.Engine {
	ginRouter := gin.Default()
	ginRouter.Use(SaveServices(service))
	v1 := ginRouter.Group("/api/v1")

	v1.GET("/user/list", handler.GetUserList)
	v1.POST("/user/login", handler.Login)
	v1.POST("/user/save-password", handler.SavePassword)
	v1.POST("/user/save", handler.Save)
	v1.POST("/user/delete", handler.DeleteUser)
	v1.GET("/user/logout", handler.Logout)
	return ginRouter
}

//SaveServices : 将服务实例存放到gin中
func SaveServices(service []interface{}) gin.HandlerFunc {
	return func(context *gin.Context) {
		//将实例存到gin.Keys里
		context.Keys = make(map[string]interface{})
		context.Keys[public.UserServiceName] = service[0]
		context.Next()
	}
}

// Cors : 跨域配置
func Cors() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Cookie", "Authorization"}
	config.AllowOrigins = []string{"*"} //http://localhost:8081
	config.AllowCredentials = true
	return cors.New(config)
}
