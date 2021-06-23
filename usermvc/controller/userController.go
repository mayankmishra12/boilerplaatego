package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
	"usermvc/model"
)

func (ctr *controller) CreateUser(c *gin.Context) {
	var user model.UserResquest
	log.Println("entering readFirst")
   if err :=c.BindJSON(&user);err != nil {
	   zap.S().Error("error while marshalling json ", err.Error())
   }

	res, err := ctr.userSvc.CreateUser(context.Background(), user)
	if err != nil {
		c.JSON(200, err.Error())
	}
	c.JSON(200, res)
}

