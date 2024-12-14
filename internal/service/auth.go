package service

import (
	"github.com/kossadda/wallet-exchanger/internal/repository"
	"github.com/kossadda/wallet-exchanger/internal/util"
	"github.com/kossadda/wallet-exchanger/model"
)

type AuthService struct {
	repo repository.Repository
}

func NewAuthService(repo repository.Repository) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (s *AuthService) CreateUser(usr model.User) error {
	usr.Password = util.GenerateHash(usr.Password, usr.Username)

	return s.repo.CreateUser(usr)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, util.GenerateHash(password, username))
	if err != nil {
		return "", err
	}

	return util.GenerateToken(&user)
}

func (s *AuthService) ParseToken(access string) (int, error) {
	return util.ParseToken(access)
}
