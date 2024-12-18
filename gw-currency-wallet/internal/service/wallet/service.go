package wallet

import (
	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/internal/model"
	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/internal/storage"
)

type MainAPI interface {
	GetBalance(userId int) (*model.Currency, error)
	Deposit(dep *model.Operation) (float64, error)
	Withdraw(with *model.Operation) (float64, error)
}

type Wallet struct {
	MainAPI
}

func New(repo *storage.Storage) *Wallet {
	return &Wallet{
		MainAPI: newService(repo),
	}
}
