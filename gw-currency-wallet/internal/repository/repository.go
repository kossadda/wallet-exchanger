package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/model"
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

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthDB(db),
		Wallet:        NewWalletDB(db),
	}
}