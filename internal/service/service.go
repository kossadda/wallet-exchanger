package service

import (
	"github.com/kossadda/wallet-exchanger/internal/repository"
	"github.com/kossadda/wallet-exchanger/model"
)

type Authorization interface {
	CreateUser(usr model.User) error
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Wallet interface {
	GetBalance(userId int) (*model.Currency, error)
	DepositSum(dep *model.Deposit) error
}

type Service struct {
	Authorization
	Wallet
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(*repos),
		Wallet:        NewWalletService(*repos),
	}
}
