package service

import (
	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/internal/model"
	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/internal/storage"
)

type Authorization interface {
	CreateUser(usr model.User) error
	GenerateToken(username, password, tokenTTL string) (string, error)
	ParseToken(token string) (int, error)
}

type Wallet interface {
	GetBalance(userId int) (*model.Currency, error)
	Deposit(dep *model.Operation) error
	Withdraw(with *model.Operation) error
}

type Service struct {
	Authorization
	Wallet
}

func NewService(repos *storage.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(*repos),
		Wallet:        NewWalletService(*repos),
	}
}
