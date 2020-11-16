package middleware

import (
	"course/public"
	"github.com/gin-gonic/gin"
)

//InitMiddleware 接收实例并存到gin.Keys里
func InitMiddleware(service []interface{}) gin.HandlerFunc {
	return func(context *gin.Context) {
		//将实例存到gin.Keys里
		context.Keys = make(map[string]interface{})
		context.Keys[public.UserServiceName] = service[0]
		context.Next()
	}
}
