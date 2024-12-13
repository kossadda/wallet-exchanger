package service

import (
	"crypto/sha1"
	"fmt"
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
	usr.Password = generateHash(usr.Password, usr.Username)

	return s.repo.CreateUser(usr)
}

func (s *AuthService) Login(usr model.User) (int, error) {
	return s.repo.Login(usr)
}

func generateHash(password string, salt string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
