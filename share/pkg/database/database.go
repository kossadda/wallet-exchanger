package database

import (
	"github.com/jmoiron/sqlx"
	"github.com/kossadda/wallet-exchanger/share/pkg/configs"
	"github.com/kossadda/wallet-exchanger/share/pkg/database/postgres"
)

const (
	UserTable     = "users"
	WalletTable   = "wallets"
	CurrencyTable = "currency"
)

type DataBase interface {
	Transaction(fn func(tx *sqlx.Tx) error) error
	Close() error
}

func NewPostgres(cfg *configs.ConfigDB) (DataBase, error) {
	return postgres.New(cfg)
}
