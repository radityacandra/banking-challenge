package repository

import (
	"context"
	"database/sql"

	"github.com/radityacandra/banking-challenge/internal/application/user-account/types"
)

func (r *Repository) BeginTransaction(ctx context.Context) (context.Context, error) {
	tx, err := r.Db.BeginTxx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}

	return context.WithValue(ctx, types.TrxKey, tx), nil
}
