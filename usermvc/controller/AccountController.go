package controller

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"usermvc/model"
)

const (

)

func (ctr *controller) InsertAccountDetails(c *gin.Context) {
	var accountDetailsReq *model.AccountDetailsRequest
	fmt.Println("getting error is the best")
	if err := c.ShouldBindJSON(&accountDetailsReq); err != nil {
		zap.S().Error("not able parse request", err.Error())
		c.JSON(200, err.Error())
		return
	}
	res, err := ctr.accountSvc.InsertAccountDetails(context.Background(), accountDetailsReq)

	if err != nil {
		zap.S().Error("not able parse request", err.Error())
		c.JSON(200, err.Error())
		return
	}
	c.JSON(200, res)
}
