package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/kossadda/wallet-exchanger/model"
)

type AuthDB struct {
	db *sqlx.DB
}

func NewAuthDB(db *sqlx.DB) *AuthDB {
	return &AuthDB{
		db: db,
	}
}

func (s *AuthDB) CreateUser(usr model.User) (int, error) {
	return 123, nil
}

func (s *AuthDB) Login(usr model.User) (int, error) {
	return 432, nil
}
