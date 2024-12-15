package database

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/kossadda/wallet-exchanger/share/configs"
	"github.com/kossadda/wallet-exchanger/share/database/postgres"
)

const (
	UserTable   = "users"
	WalletTable = "wallets"
)

type DataBase interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sqlx.Rows, error)
	QueryRow(query string, args ...interface{}) *sqlx.Row
	Transaction(fn func(tx *sqlx.Tx) error) error
	Close() error
}

func NewPostgres(cfg *configs.Config) (DataBase, error) {
	return postgres.New(cfg)
}
