package wallet

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/internal/model"
	"github.com/kossadda/wallet-exchanger/share/pkg/database"
)

var SupportCurrency = map[string]struct{}{"USD": {}, "RUB": {}, "EUR": {}}

type storage struct {
	db database.DataBase
}

func newStorage(db database.DataBase) *storage {
	return &storage{
		db: db,
	}
}

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
