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

	admin := v1.Group("/admin")
	{
		user := admin.Group("/user")
		{
			user.GET("/list", handler.GetUserList)
			user.POST("/login", handler.Login)
			user.POST("/save-password", handler.SavePassword)
			user.POST("/save", handler.Save)
			user.POST("/delete", handler.DeleteUser)
			user.GET("/logout", handler.Logout)
		}
		admin.GET("/captcha/image-code", handler.GetCaptcha)
		resource := admin.Group("/resource")
		{
			resource.GET("/load-tree", handler.LoadTree)
			resource.POST("/save", handler.SaveJson)
			resource.DELETE("/delete", handler.Delete)
		}
		role := admin.Group("/role")
		{
			role.POST("/list", handler.RoleList)
			role.POST("/save", handler.SaveRole)
			role.DELETE("/delete", handler.DeleteRole)
			role.POST("/save-resource", handler.SaveRoleResource)
			role.GET("/list-resource", handler.ListRoleResource)
			role.POST("/save-user", handler.SaveRoleUser)
			role.GET("/list-user", handler.ListRoleUser)
		}
	}

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
