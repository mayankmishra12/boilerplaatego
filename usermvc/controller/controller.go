package controller

import (
	"usermvc/service/account"
	"usermvc/service/user"
)

type controller struct {
	userSvc    user.Service
	accountSvc account.Service
}

func NewController() *controller {
	return &controller{userSvc: user.NewuserService(), accountSvc: account.NewAccountService()}
}
