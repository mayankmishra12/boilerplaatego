package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"usermvc/model"
)

func (ctr *controller) CreateUser(c *gin.Context) {
	var user model.UserResquest
	c.BindJSON(&user)
	res, err := ctr.userSvc.CreateUser(context.Background(), user)
	if err != nil {
		c.JSON(200, err.Error())
	}
	c.JSON(200, res)
}
