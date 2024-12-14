package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/kossadda/wallet-exchanger/model"
)

type WalletDB struct {
	db *sqlx.DB
}

func NewWalletDB(db *sqlx.DB) *WalletDB {
	return &WalletDB{
		db: db,
	}
}

func (w *WalletDB) GetBalance(userId int) (model.BalanceCurrency, error) {
	return model.BalanceCurrency{}, nil
}
