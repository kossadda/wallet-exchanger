// Package wallet provides functionality related to user wallets, including balance retrieval, deposits, and withdrawals.
package wallet

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/internal/model"
	"github.com/kossadda/wallet-exchanger/share/pkg/database"
)

// SupportCurrency defines a set of supported currencies for deposit and withdrawal operations.
var SupportCurrency = map[string]struct{}{"USD": {}, "RUB": {}, "EUR": {}}

// storage represents the underlying storage layer for wallet operations, interacting with the database.
type storage struct {
	db database.DataBase
}

// newStorage creates and returns a new instance of storage.
func newStorage(db database.DataBase) *storage {
	return &storage{
		db: db,
	}
}

// GetBalance retrieves the user's balance in multiple currencies (USD, RUB, EUR).
// It returns a model.Currency object containing the balances for each supported currency.
func (w *storage) GetBalance(userId int) (*model.Currency, error) {
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

// Deposit adds a specified amount of currency to the user's wallet.
// The currency type must be one of the supported currencies (USD, RUB, EUR).
// It returns the updated balance for the specified currency.
func (w *storage) Deposit(dep *model.Operation) (float64, error) {
	if _, exist := SupportCurrency[dep.Currency]; !exist {
		return 0, fmt.Errorf("currency %s not supported", dep.Currency)
	}

	cur := strings.ToLower(dep.Currency)

	var updatedBalance float64
	if err := w.db.Transaction(func(tx *sqlx.Tx) error {
		query := fmt.Sprintf("UPDATE %s SET %s = %s + $1 WHERE id = $2 RETURNING %s", database.WalletTable, cur, cur, cur)

		if err := tx.QueryRow(query, dep.Amount, dep.UserId).Scan(&updatedBalance); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return 0, err
	}

	return updatedBalance, nil
}

// Withdraw removes a specified amount of currency from the user's wallet.
// The currency type must be one of the supported currencies (USD, RUB, EUR).
// It returns the updated balance after the withdrawal or an error if there are insufficient funds.
func (w *storage) Withdraw(with *model.Operation) (float64, error) {
	if _, exist := SupportCurrency[with.Currency]; !exist {
		return 0, fmt.Errorf("currency %s not supported", with.Currency)
	}

	var currentBalance float64
	if err := w.db.Transaction(func(tx *sqlx.Tx) error {
		selectQuery := fmt.Sprintf("SELECT %s FROM %s WHERE id = $1", strings.ToLower(with.Currency), database.WalletTable)
		updateQuery := fmt.Sprintf("UPDATE %s SET %s = %s - $1 WHERE id = $2", database.WalletTable, strings.ToLower(with.Currency), strings.ToLower(with.Currency))

		if err := tx.QueryRow(selectQuery, with.UserId).Scan(&currentBalance); err != nil {
			return err
		}

		if currentBalance < with.Amount {
			return fmt.Errorf("insufficient funds or invalid amount")
		}

		if _, err := tx.Exec(updateQuery, with.Amount, with.UserId); err != nil {
			return err
		}

		if err := tx.QueryRow(selectQuery, with.UserId).Scan(&currentBalance); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return 0, err
	}

	return currentBalance, nil
}
