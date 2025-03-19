package repository

import (
	"context"

	"github.com/radityacandra/banking-challenge/internal/application/user-account/model"
)

func (r *Repository) FindUserAccountByAccountNo(ctx context.Context, accountNo string) (*model.UserAccount, error) {
	row := r.Db.QueryRowxContext(ctx, `
		SELECT
			*
		FROM public.user_accounts
		WHERE
			account_number = $1
	`, accountNo)

	var userAccount model.UserAccount
	if err := row.StructScan(&userAccount); err != nil {
		return nil, err
	}

	return &userAccount, nil
}
