package service

import (
	"github.com/kossadda/wallet-exchanger/internal/repository"
	"github.com/kossadda/wallet-exchanger/model"
)

type Authorization interface {
	CreateUser(usr model.User) (int, error)
	Login(usr model.User) (int, error)
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(*repos),
	}
}
