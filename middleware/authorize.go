package middleware

import (
	"github.com/gin-gonic/gin"
)

func AuthorizationHook() gin.HandlerFunc {
	return func(c *gin.Context) {
		// username := c.GetHeader("username")
		// pin := c.GetHeader("pin")
		// common.LogInfo(nil, fmt.Sprintf("username: %s, pin: %s", username, pin))
		c.Next()
	}
}
