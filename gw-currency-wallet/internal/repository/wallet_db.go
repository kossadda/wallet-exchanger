package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/model"
	"github.com/kossadda/wallet-exchanger/share/database"
)

type WalletDB struct {
	db database.DataBase
}

func NewWalletDB(db database.DataBase) *WalletDB {
	return &WalletDB{
		db: db,
	}
}

func (w *WalletDB) GetBalance(userId int) (*model.Currency, error) {
	var balance model.Currency

	if err := w.db.Transaction(func(tx *sqlx.Tx) error {
		query := fmt.Sprintf("SELECT usd, rub, eur FROM %s WHERE id=$1", database.WalletTable)
		err := tx.Get(&balance, query, userId)
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return &balance, nil
}

func (w *WalletDB) Deposit(dep *model.Operation) error {
	return w.db.Transaction(func(tx *sqlx.Tx) error {
		var query string
		switch dep.Currency {
		case "USD":
			query = fmt.Sprintf("UPDATE %s SET usd = usd + $1 WHERE id = $2", database.WalletTable)
		case "RUB":
			query = fmt.Sprintf("UPDATE %s SET rub = rub + $1 WHERE id = $2", database.WalletTable)
		case "EUR":
			query = fmt.Sprintf("UPDATE %s SET eur = eur + $1 WHERE id = $2", database.WalletTable)
		default:
			return fmt.Errorf("unsupported currency: %s", dep.Currency)
		}

		if _, err := tx.Exec(query, dep.Amount, dep.UserId); err != nil {
			return err
		}

		return nil
	})
}

func (w *WalletDB) Withdraw(with *model.Operation) error {
	return w.db.Transaction(func(tx *sqlx.Tx) error {
		var currentBalance float64
		var selectQuery, updateQuery string

		switch with.Currency {
		case "USD":
			selectQuery = fmt.Sprintf("SELECT usd FROM %s WHERE id = $1", database.WalletTable)
			updateQuery = fmt.Sprintf("UPDATE %s SET usd = usd - $1 WHERE id = $2", database.WalletTable)
		case "RUB":
			selectQuery = fmt.Sprintf("SELECT rub FROM %s WHERE id = $1", database.WalletTable)
			updateQuery = fmt.Sprintf("UPDATE %s SET rub = rub - $1 WHERE id = $2", database.WalletTable)
		case "EUR":
			selectQuery = fmt.Sprintf("SELECT eur FROM %s WHERE id = $1", database.WalletTable)
			updateQuery = fmt.Sprintf("UPDATE %s SET eur = eur - $1 WHERE id = $2", database.WalletTable)
		default:
			return fmt.Errorf("unsupported currency: %s", with.Currency)
		}

		if err := tx.QueryRow(selectQuery, with.UserId).Scan(&currentBalance); err != nil {
			return err
		}

		if currentBalance < with.Amount {
			return fmt.Errorf("insufficient funds or invalid amount")
		}

		if _, err := tx.Exec(updateQuery, with.Amount, with.UserId); err != nil {
			return err
		}

		return nil
	})
}
