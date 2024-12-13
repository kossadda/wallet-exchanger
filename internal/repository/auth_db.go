package repository

import (
	"database/sql"
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

func (s *AuthDB) CreateUser(usr model.User) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	query := fmt.Sprintf("INSERT INTO %s (username, password_hash, email) VALUES ($1, $2, $3)", model.UserTable)
	row := s.db.QueryRow(query, usr.Username, usr.Password, usr.Email)
	if err := row.Scan(); err != nil {
		if err != sql.ErrNoRows {
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (s *AuthDB) Login(usr model.User) error {
	return nil
}
