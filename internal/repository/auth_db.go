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

	query := fmt.Sprintf("INSERT INTO %s (username, password_hash, email) VALUES ($1, $2, $3) RETURNING id", model.UserTable)
	var userID int
	err = tx.QueryRow(query, usr.Username, usr.Password, usr.Email).Scan(&userID)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	query = fmt.Sprintf("INSERT INTO %s (user_id) VALUES ($1)", model.WalletTable)
	_, err = tx.Exec(query, userID)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (s *AuthDB) GetUser(username, password string) (*model.User, error) {
	var user model.User
	tx, err := s.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", model.UserTable)
	err = s.db.Get(&user, query, username, password)

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &user, err
}
