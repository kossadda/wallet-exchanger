package service

import (
	"time"

	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/internal/model"
	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/internal/storage"
	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/internal/util"
)

const (
	defaultTokenTTL = time.Hour * 24
)

type AuthService struct {
	repo storage.Repository
}

func NewAuthService(repo storage.Repository) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (s *AuthService) CreateUser(usr model.User) error {
	usr.Password = util.GenerateHash(usr.Password, usr.Username)

	return s.repo.CreateUser(usr)
}

func (s *AuthService) GenerateToken(username, password, tokenTTL string) (string, error) {
	user, err := s.repo.GetUser(username, util.GenerateHash(password, username))
	if err != nil {
		return "", err
	}

	ttl, err := time.ParseDuration(tokenTTL)
	if err != nil {
		ttl = defaultTokenTTL
	}

	return util.GenerateToken(user, ttl)
}

func (s *AuthService) ParseToken(access string) (int, error) {
	return util.ParseToken(access)
}
