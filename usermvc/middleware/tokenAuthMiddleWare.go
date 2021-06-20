package middleware

import "github.com/gin-gonic/gin"

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//do token validation
		c.Next()
	}
}