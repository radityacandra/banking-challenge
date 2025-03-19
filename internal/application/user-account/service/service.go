package service

import (
	"context"

	"github.com/radityacandra/banking-challenge/internal/application/user-account/repository"
	"github.com/radityacandra/banking-challenge/internal/application/user-account/types"
)

type IService interface {
	GetBalance(ctx context.Context, accountNo string) (types.GetBalanceOutput, error)
	StoreCash(ctx context.Context, input types.TransactionInput) (types.GetBalanceOutput, error)
	WithdrawCash(ctx context.Context, input types.TransactionInput) (types.GetBalanceOutput, error)
}

type Service struct {
	Repository repository.IRepository
}

func NewService(repository repository.IRepository) *Service {
	return &Service{
		Repository: repository,
	}
}
