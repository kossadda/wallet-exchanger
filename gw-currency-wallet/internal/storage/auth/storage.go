package auth

import (
	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/internal/model"
	"github.com/kossadda/wallet-exchanger/share/pkg/database"
)

// MainAPI defines the methods related to authentication, such as creating and retrieving users.
type MainAPI interface {
	// CreateUser creates a new user
	CreateUser(user model.User) error

	// GetUser retrieves a user by username and password
	GetUser(username, password string) (*model.User, error)
}

// Auth is the main struct for handling authentication logic and implements the MainAPI interface.
type Auth struct {
	MainAPI
}

// New creates and returns a new instance of Auth.
func New(db database.DataBase) *Auth {
	return &Auth{
		MainAPI: newStorage(db),
	}
}
