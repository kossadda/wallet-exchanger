package auth

import (
	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/internal/model"
	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/internal/storage"
)

// MainAPI defines the core operations related to user authentication, such as creating users, generating tokens, and parsing tokens.
type MainAPI interface {
	// CreateUser creates a new user in the storage after hashing their password.
	CreateUser(usr model.User) error

	// GenerateToken generates an authentication token for a user.
	GenerateToken(username, password, tokenTTL string) (string, error)

	// ParseToken parses an access token and returns the user ID.
	ParseToken(token string) (int, error)
}

// Auth is a service that implements the MainAPI interface, providing authentication functionalities.
type Auth struct {
	MainAPI
}

// New creates and returns a new Auth instance using the provided storage repository.
func New(repo *storage.Storage) *Auth {
	return &Auth{
		MainAPI: newService(repo),
	}
}
