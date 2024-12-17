package wallet

import (
	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/internal/model"
	"github.com/kossadda/wallet-exchanger/share/pkg/database"
)

type MainAPI interface {
	GetBalance(userId int) (*model.Currency, error)
	Deposit(dep *model.Operation) error
	Withdraw(with *model.Operation) error
}

type Wallet struct {
	MainAPI
}

func New(db database.DataBase) *Wallet {
	return &Wallet{
		MainAPI: newDatabase(db),
	}
}
