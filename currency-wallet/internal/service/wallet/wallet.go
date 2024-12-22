// Package wallet provides functionality related to user wallets, including balance retrieval, deposits, and withdrawals.
package wallet

import (
	"github.com/kossadda/wallet-exchanger/currency-wallet/internal/model"
	"github.com/kossadda/wallet-exchanger/currency-wallet/internal/storage"
)

// service handles wallet operations such as retrieving balance, depositing, and withdrawing funds.
type service struct {
	repo *storage.Storage
}

// newService creates and returns a new instance of service.
func newService(repo *storage.Storage) *service {
	return &service{
		repo: repo,
	}
}

// GetBalance retrieves the current balance of a user from the storage.
func (w *service) GetBalance(userId int) (*model.Currency, error) {
	return w.repo.GetBalance(userId)
}

// Deposit adds a specified amount to the user's wallet and returns the updated balance.
func (w *service) Deposit(dep *model.Operation) (float64, error) {
	return w.repo.Deposit(dep)
}

// Withdraw removes a specified amount from the user's wallet and returns the updated balance.
func (w *service) Withdraw(dep *model.Operation) (float64, error) {
	return w.repo.Withdraw(dep)
}
