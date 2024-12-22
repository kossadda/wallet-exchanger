package wallet

import (
	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/internal/model"
	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/internal/storage"
)

// MainAPI defines the core wallet operations: retrieving balance, depositing funds, and withdrawing funds.
type MainAPI interface {
	// GetBalance retrieves the balance for a user in multiple currencies.
	GetBalance(userId int) (*model.Currency, error)

	// Deposit adds funds to the user's wallet.
	Deposit(dep *model.Operation) (float64, error)

	// Withdraw removes funds from the user's wallet.
	Withdraw(with *model.Operation) (float64, error)
}

// Wallet represents a wallet service that implements the MainAPI interface.
type Wallet struct {
	MainAPI
}

// New creates and returns a new instance of Wallet with the given repository.
func New(repo *storage.Storage) *Wallet {
	return &Wallet{
		MainAPI: newService(repo),
	}
}
