package controller

import "usermvc/service"

type controller struct {
	userSvc  service.UserService
}
func NewController() *controller {
return &controller{userSvc: service.NewuserService()}
}
