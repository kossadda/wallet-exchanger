// Package auth handles user authentication functionalities such as user creation, token generation, and token parsing.
package auth

import (
	"time"

	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/internal/model"
	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/internal/storage"
	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/internal/util"
	"github.com/kossadda/wallet-exchanger/share/pkg/configs"
)

// service provides the authentication functionality by interacting with the storage layer.
type service struct {
	repo *storage.Storage
}

// newService creates and returns a new instance of service.
func newService(repo *storage.Storage) *service {
	return &service{
		repo: repo,
	}
}

// CreateUser hashes the user's password and saves the user to the storage.
func (s *service) CreateUser(usr model.User) error {
	usr.Password = util.GenerateHash(usr.Password, usr.Username)

	return s.repo.CreateUser(usr)
}

// GenerateToken generates an authentication token for a user, using their credentials and an optional token TTL.
func (s *service) GenerateToken(username, password, tokenTTL string) (string, error) {
	user, err := s.repo.GetUser(username, util.GenerateHash(password, username))
	if err != nil {
		return "", err
	}

	ttl, err := time.ParseDuration(tokenTTL)
	if err != nil {
		ttl = configs.DefaultTokenExpire
	}

	return util.GenerateToken(user, ttl)
}

// ParseToken parses an access token and returns the user ID embedded within.
func (s *service) ParseToken(access string) (int, error) {
	return util.ParseToken(access)
}
