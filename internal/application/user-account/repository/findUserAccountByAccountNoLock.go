package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/radityacandra/banking-challenge/internal/application/user-account/model"
	"github.com/radityacandra/banking-challenge/internal/application/user-account/types"
)

func (r *Repository) FindUserAccountByAccountNoLock(ctx context.Context, accountNo string) (*model.UserAccount, error) {
	tx, ok := ctx.Value(types.TrxKey).(*sqlx.Tx)
	if !ok {
		return nil, types.ErrFailedToAcquireTx
	}

	row := tx.QueryRowxContext(ctx, `
		SELECT
			*
		FROM public.user_accounts
		WHERE
			account_number = $1
		FOR UPDATE
	`, accountNo)

	var userAccount model.UserAccount
	if err := row.StructScan(&userAccount); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = errors.Join(err, types.ErrUserAccountNotFound)
		}

		return nil, errors.Join(err, tx.Rollback())
	}

	return &userAccount, nil
}
