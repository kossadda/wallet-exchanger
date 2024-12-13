package repository

import (
	"fmt"

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
	var id int
	query := fmt.Sprintf("INSERT INTO %s (username, password_hash, email) VALUES ($1, $2, $3) RETURNING id", model.UserTable)

	row := s.db.QueryRow(query, usr.Username, usr.Password, usr.Email)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (s *AuthDB) Login(usr model.User) (int, error) {
	return 432, nil
}
