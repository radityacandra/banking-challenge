package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/radityacandra/banking-challenge/internal/application/user-account/model"
	"github.com/radityacandra/banking-challenge/internal/application/user-account/types"
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
		if errors.Is(err, sql.ErrNoRows) {
			err = errors.Join(err, types.ErrUserAccountNotFound)
		}

		return nil, err
	}

	return &userAccount, nil
}
