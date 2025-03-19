package repository

import (
	"context"
	"errors"

	"github.com/jmoiron/sqlx"
	userAccount "github.com/radityacandra/banking-challenge/internal/application/user-account/model"
	"github.com/radityacandra/banking-challenge/internal/application/user/model"
)

func (r *Repository) InsertUser(ctx context.Context, user model.User) error {
	tx, err := r.Db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	query, args, _ := sqlx.BindNamed(sqlx.DOLLAR, `
		INSERT INTO public.users(id, name, phone_number, identity_number, created_at, created_by)
		VALUES(:id, :name, :phone_number, :identity_number, :created_at, :created_by)`, &user)

	_, err = tx.ExecContext(ctx, query, args...)
	if err != nil {
		return errors.Join(err, tx.Rollback())
	}

	userAccount := user.Accounts[0].(*userAccount.UserAccount)

	query, args, _ = sqlx.BindNamed(sqlx.DOLLAR, `
		INSERT INTO public.user_accounts(id, user_id, account_number, total_balance, created_at, created_by)
		VALUES(:id, :user_id, :account_number, :total_balance, :created_at, :created_by)
	`, &userAccount)

	_, err = tx.ExecContext(ctx, query, args...)
	if err != nil {
		return errors.Join(err, tx.Rollback())
	}

	return tx.Commit()
}
