package service

import (
	"github.com/kossadda/wallet-exchanger/internal/repository"
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

func (s *AuthService) CreateUser(usr model.User) (int, error) {
	return s.repo.CreateUser(usr)
}

func (s *AuthService) Login(usr model.User) (int, error) {
	return s.repo.Login(usr)
}
