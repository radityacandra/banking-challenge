package service

import (
	"context"

	"github.com/google/uuid"
	userAccount "github.com/radityacandra/banking-challenge/internal/application/user-account/model"
	"github.com/radityacandra/banking-challenge/internal/application/user/model"
	"github.com/radityacandra/banking-challenge/internal/application/user/types"
)

func (s *Service) RegisterUser(ctx context.Context, input types.RegisterUserInput) (types.RegisterUserOutput, error) {
	userId := uuid.NewString()
	accountId := uuid.NewString()
	account := userAccount.NewUserAccount(accountId, userId, uuid.NewString()) // TODO: generate account no
	user := model.NewUser(userId, input.Name, input.PhoneNo, input.IdentityNo, account)

	// insert to database
	err := s.Repository.InsertUser(ctx, *user)
	if err != nil {
		return types.RegisterUserOutput{}, err
	}

	return types.RegisterUserOutput{
		UserId:        userId,
		BankAccountNo: account.AccountNo,
	}, nil
}
