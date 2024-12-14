package repository

import (
	"fmt"
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

func (w *WalletDB) GetBalance(userId int) (*model.BalanceCurrency, error) {
	var balance model.Currency

	tx, err := w.db.Beginx()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	query := fmt.Sprintf("SELECT usd, rub, eur FROM %s WHERE id=$1", model.WalletTable)
	err = tx.Get(&balance, query, userId)
	if err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return &model.BalanceCurrency{Balance: balance}, nil
}
