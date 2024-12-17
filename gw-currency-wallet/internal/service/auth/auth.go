package auth

import (
	"time"

	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/internal/model"
	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/internal/storage"
	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/internal/util"
	"github.com/kossadda/wallet-exchanger/share/pkg/configs"
)

type service struct {
	repo *storage.Storage
}

func newService(repo *storage.Storage) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateUser(usr model.User) error {
	usr.Password = util.GenerateHash(usr.Password, usr.Username)

	return s.repo.CreateUser(usr)
}

func (s *service) GenerateToken(username, password, tokenTTL string) (string, error) {
	user, err := s.repo.GetUser(username, util.GenerateHash(password, username))
	if err != nil {
		return "", err
	}

	ttl, err := time.ParseDuration(tokenTTL)
	if err != nil {
		ttl = configs.DefaultTokenTTL
	}

	return util.GenerateToken(user, ttl)
}

func (s *service) ParseToken(access string) (int, error) {
	return util.ParseToken(access)
}
