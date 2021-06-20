package service

import (
	"context"
	"usermvc/entity"
	"usermvc/model"
	"usermvc/repositories/userrepo"
)


type UserService interface {
	CreateUser(ctx context.Context, resquest model.UserResquest) (*model.UserResponse, error)
}

type userService struct {
	userRepo userrepo.UserRepo
}
func NewuserService () *userService {
	return &userService{
		userRepo: userrepo.NewUserRepo(),
	}
}

func (s userService)CreateUser(ctx context.Context, resquest model.UserResquest) (*model.UserResponse, error)  {
	user := entity.User{
		Email:        resquest.Email,
		FirstName:    resquest.FirstName,
		LastName:     resquest.LastName,
		Entitlements: resquest.Entitlements,
	}
  if  err :=  s.userRepo.Create(ctx, user);err != nil{
  	return nil,err
  }
	return &model.UserResponse{Status: 232},nil
}