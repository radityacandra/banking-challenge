package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/radityacandra/banking-challenge/internal/application/user/model"
)

type IRepository interface {
	InsertUser(ctx context.Context, user model.User) error
}

type Repository struct {
	Db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Db: db,
	}
}
