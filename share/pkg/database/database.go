package database

import (
	"github.com/jmoiron/sqlx"
	"github.com/kossadda/wallet-exchanger/share/pkg/configs"
	"github.com/kossadda/wallet-exchanger/share/pkg/database/postgres"
)

// Constants for table names.
const (
	UserTable     = "users"
	WalletTable   = "wallets"
	CurrencyTable = "currency"
)

// DataBase defines the methods for interacting with a database.
type DataBase interface {
	Transaction(fn func(tx *sqlx.Tx) error) error
	Close() error
}

// NewPostgres creates a new PostgreSQL database connection.
func NewPostgres(cfg *configs.ConfigDB) (DataBase, error) {
	return postgres.New(cfg)
}
