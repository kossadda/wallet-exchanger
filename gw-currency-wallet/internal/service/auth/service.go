package auth

import (
	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/internal/model"
	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/internal/storage"
)

type MainAPI interface {
	CreateUser(usr model.User) error
	GenerateToken(username, password, tokenTTL string) (string, error)
	ParseToken(token string) (int, error)
}

type Auth struct {
	MainAPI
}

func New(repo *storage.Storage) *Auth {
	return &Auth{
		MainAPI: newService(repo),
	}
}
