package user

import (
	"context"
	"usermvc/entity"
	"usermvc/model"
	"usermvc/repositories/userrepo"
)

type Service interface {
	CreateUser(ctx context.Context, resquest model.UserResquest) (*model.UserResponse, error)
}

type service struct {
	userRepo userrepo.UserRepo
}

func NewuserService() *service {
	return &service{
		userRepo: userrepo.NewUserRepo(),
	}
}

func (s service) CreateUser(ctx context.Context, resquest model.UserResquest) (*model.UserResponse, error) {
	user := entity.User{
		Email:        resquest.Email,
		FirstName:    resquest.FirstName,
		LastName:     resquest.LastName,
		Entitlements: resquest.Entitlements,
	}
	if err := s.userRepo.Create(ctx, user); err != nil {
		return nil, err
	}
	return &model.UserResponse{Status: 232}, nil
}
