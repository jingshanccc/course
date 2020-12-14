package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//JWT: OAUTH校验JWT TOKEN
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := AuthServer.ValidationBearerToken(c.Request)
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}
		c.Next()
	}
}
