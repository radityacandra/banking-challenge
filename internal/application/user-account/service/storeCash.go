package service

import (
	"context"

	"github.com/radityacandra/banking-challenge/internal/application/user-account/types"
)

func (s *Service) StoreCash(ctx context.Context, input types.TransactionInput) (types.GetBalanceOutput, error) {
	txCtx, err := s.Repository.BeginTransaction(ctx)
	if err != nil {
		return types.GetBalanceOutput{}, err
	}

	userAccount, err := s.Repository.FindUserAccountByAccountNoLock(txCtx, input.AccountNo)
	if err != nil {
		return types.GetBalanceOutput{}, err
	}

	latestHistory := userAccount.TransactionCredit(input.Amount)

	err = s.Repository.SaveTransaction(txCtx, *userAccount, *latestHistory)
	if err != nil {
		return types.GetBalanceOutput{}, err
	}

	return types.GetBalanceOutput{
		Balance: userAccount.TotalBalance,
	}, err
}
