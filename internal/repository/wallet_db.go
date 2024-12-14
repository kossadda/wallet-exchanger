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

func (w *WalletDB) GetBalance(userId int) (*model.Currency, error) {
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

	return &balance, nil
}

func (w *WalletDB) DepositSum(dep *model.Deposit) error {
	tx, err := w.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var query string
	switch dep.Currency {
	case "USD":
		query = fmt.Sprintf("UPDATE %s SET usd = usd + $1 WHERE id = $2", model.WalletTable)
	case "RUB":
		query = fmt.Sprintf("UPDATE %s SET rub = rub + $1 WHERE id = $2", model.WalletTable)
	case "EUR":
		query = fmt.Sprintf("UPDATE %s SET eur = eur + $1 WHERE id = $2", model.WalletTable)
	default:
		return fmt.Errorf("unsupported currency: %s", dep.Currency)
	}

	_, err = tx.Exec(query, dep.Amount, dep.UserId)
	if err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}
