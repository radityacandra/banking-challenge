package service

import (
	"context"

	"github.com/radityacandra/banking-challenge/internal/application/user-account/types"
)

func (s *Service) GetBalance(ctx context.Context, accountNo string) (types.GetBalanceOutput, error) {
	userAccount, err := s.Repository.FindUserAccountByAccountNo(ctx, accountNo)
	if err != nil {
		return types.GetBalanceOutput{}, err
	}

	return types.GetBalanceOutput{
		AccountNo: userAccount.AccountNo,
		Balance:   userAccount.TotalBalance,
	}, nil
}
