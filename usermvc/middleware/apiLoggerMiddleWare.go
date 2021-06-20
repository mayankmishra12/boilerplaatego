package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"io/ioutil"
	"usermvc/utility/logger"
)
type key string
var(
	_requestID = "requestID"

)
func LoggerMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := GetHttpRequestContext(c)
		logger := logger.GetLoggerWithContext(ctx)
		requestParams := c.Request.URL.Query()
		params, err := json.Marshal(requestParams)
		if err != nil {
			logger.Fatal(err)
		}
		requestBody,err  := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			logger.Error("error while converting request body")
		}
        logger.Info("getting api request ","method ",c.Request.Method, " query params ", string(params), "request data ", string(requestBody))
		logger.Info("request name")
		c.Next()
	}
}

func GetHttpRequestContext (ctx *gin.Context) context.Context{
	requestId := ctx.GetHeader(_requestID)
	if requestId == "" {
		requestId = generateReuestID()
	}
	httpContext:= context.WithValue(ctx, key(_requestID), generateReuestID())
	return httpContext
}
func generateReuestID() string {
	requestID, _ := uuid.NewV4()
	fmt.Println("printing requestid", requestID)
	return requestID.String()

}