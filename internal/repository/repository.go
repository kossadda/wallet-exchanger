package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/kossadda/wallet-exchanger/model"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
	Login(usr model.User) (int, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthDB(db),
	}
}
