package account

import (
	"context"
	"encoding/json"
	"fmt"
	"usermvc/entity"
	model2 "usermvc/model"
	"usermvc/repositories/accountrepo"
	"usermvc/utility/logger"
)

type Service interface {
	GetAccountDetails(ctx context.Context, request *model2.AccountDetailsRequest) (*model2.AccountDetailsResponse, error)
	InsertAccountDetails(ctx context.Context, request *model2.AccountDetailsRequest) (*model2.AccountDetailsResponse, error)
}

type service struct {
	accountRepo accountrepo.AccountRepo
}

func NewAccountService() *service {
	return &service{
		accountRepo: accountrepo.NewAccountRepo(),
	}
}

func (s service)GetAccountDetails(ctx context.Context, request *model2.AccountDetailsRequest) (*model2.AccountDetailsResponse, error) {
	//if err := s.userRepo.Create(ctx, user); err != nil {
	//	return nil, err
	//}
	//return &model.UserRessponse{Status: 232}, nil
	return nil, nil
}

func (s service) InsertAccountDetails( ctx context.Context, request *model2.AccountDetailsRequest) (*model2.AccountDetailsResponse, error)  {
	logger := logger.GetLoggerWithContext(ctx)
	fmt.Println("reaching ith thesat services")
	res, err :=s.accountRepo.Insert(ctx,entity.AccountDetails{
		LeadId:               request.LeadId,
		Role:                 request.Role,
		ConvertLeadToAccount: request.ConvertLeadToAccount,
		Approve:              request.Approve,
		Reject:               request.Reject,
		Comments:             request.Comments,
	})
	if err!= nil {
		logger.Error("not able parse request", err.Error())
		return nil, err
	}
	e, err := json.Marshal(res.Payload)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Println(string(e))
	logger.Info("getting response from account details, ",  string(e) )
	return &model2.AccountDetailsResponse{
		StatusCode: 200,
		Payload:    res,
	},nil
}