package auth

import (
	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/internal/model"
	"github.com/kossadda/wallet-exchanger/share/pkg/database"
)

type MainAPI interface {
	CreateUser(user model.User) error
	GetUser(username, password string) (*model.User, error)
}

type Auth struct {
	MainAPI
}

func New(db database.DataBase) *Auth {
	return &Auth{
		MainAPI: newDatabase(db),
	}
}
