package routes

import (
	"github.com/gin-gonic/gin"
	"usermvc/controller"
	"usermvc/middleware"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	controller := controller.NewController()
	r := gin.Default()
	r.Use(middleware.LoggerMiddleWare())
	grp1 := r.Group("/usersvc")

	{
		grp1.POST("user", controller.CreateUser)

	}
	return r
}
