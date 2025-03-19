package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/radityacandra/banking-challenge/internal/application/user-account/model"
)

type IRepository interface {
	FindUserAccountByAccountNo(ctx context.Context, accountNo string) (*model.UserAccount, error)
	BeginTransaction(ctx context.Context) (context.Context, error)
	FindUserAccountByAccountNoLock(ctx context.Context, accountNo string) (*model.UserAccount, error)
	SaveTransaction(ctx context.Context, userAccount model.UserAccount, transactionHistory model.TransactionHistory) error
}

type Repository struct {
	Db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Db: db,
	}
}
