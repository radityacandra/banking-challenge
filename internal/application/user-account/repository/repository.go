package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/radityacandra/banking-challenge/internal/application/user-account/model"
)

type IRepository interface {
	FindUserAccountByAccountNo(ctx context.Context, accountNo string) (*model.UserAccount, error)
}

type Repository struct {
	Db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Db: db,
	}
}
