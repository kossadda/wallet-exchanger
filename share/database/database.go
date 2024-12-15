package database

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/kossadda/wallet-exchanger/share/configs"
	"github.com/kossadda/wallet-exchanger/share/database/postgres"
	_ "github.com/lib/pq"
)

const (
	UserTable   = "users"
	WalletTable = "wallets"
)

type DataBase interface {
	Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	Query(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error)
	QueryRow(ctx context.Context, query string, args ...interface{}) *sqlx.Row
	Transaction(ctx context.Context, fn func(tx *sqlx.Tx) error) error
	Close() error
}

func NewPostgres(cfg *configs.Config) (DataBase, error) {
	return postgres.New(cfg)
}
