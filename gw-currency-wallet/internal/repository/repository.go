package repository

import (
	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/model"
	"github.com/kossadda/wallet-exchanger/share/database"
)

type Authorization interface {
	CreateUser(user model.User) error
	GetUser(username, password string) (*model.User, error)
}

type Wallet interface {
	GetBalance(userId int) (*model.Currency, error)
	Deposit(dep *model.Operation) error
	Withdraw(with *model.Operation) error
}

type Repository struct {
	Authorization
	Wallet
}

func NewRepository(db database.DataBase) *Repository {
	return &Repository{
		Authorization: NewAuthDB(db),
		Wallet:        NewWalletDB(db),
	}
}
