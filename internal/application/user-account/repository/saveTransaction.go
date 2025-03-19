package repository

import (
	"context"
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/radityacandra/banking-challenge/internal/application/user-account/model"
	"github.com/radityacandra/banking-challenge/internal/application/user-account/types"
)

func (r *Repository) SaveTransaction(ctx context.Context, userAccount model.UserAccount, transactionHistory model.TransactionHistory) error {
	tx, ok := ctx.Value(types.TrxKey).(*sqlx.Tx)
	if !ok {
		return types.ErrFailedToAcquireTx
	}

	_, err := tx.NamedExecContext(ctx, `
		UPDATE public.user_accounts
		SET
			total_balance = :total_balance,
			updated_at = :updated_at,
			updated_by = :updated_by
		WHERE
			account_number = :account_number
	`, &userAccount)
	if err != nil {
		return errors.Join(err, tx.Rollback())
	}

	_, err = tx.NamedExecContext(ctx, `
		INSERT INTO public.transactions(id, user_account_id, transaction_type, amount, created_at, created_by)
		VALUES(:id, :user_account_id, :transaction_type, :amount, :created_at, :created_by)
	`, &transactionHistory)
	if err != nil {
		return errors.Join(err, tx.Rollback())
	}

	return tx.Commit()
}
