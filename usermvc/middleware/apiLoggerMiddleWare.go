package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"usermvc/utility/logger"
)

func LoggerMiddleWare()  gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.SetTraceID(c)
		logger  := logger.GetLoggerWithContext(c)
		logger.Error("teri maa bsdk")
		zap.S().Info("Data encoded => ", c.Request.Body)
		//do token validation
		c.Next()
	}
}
