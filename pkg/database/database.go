package database

import (
	"context"
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

func Init(ctx context.Context, dbDsn string) (*sqlx.DB, error) {
	db, err := sql.Open("pgx", dbDsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return sqlx.NewDb(db, "pgx"), nil
}
