package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/model"
	"github.com/kossadda/wallet-exchanger/share/database"
)

type AuthDB struct {
	db database.DataBase
}

func NewAuthDB(db database.DataBase) *AuthDB {
	return &AuthDB{
		db: db,
	}
}

func (s *AuthDB) CreateUser(usr model.User) error {
	return s.db.Transaction(func(tx *sqlx.Tx) error {
		query := fmt.Sprintf("INSERT INTO %s (username, password_hash, email) VALUES ($1, $2, $3) RETURNING id", database.UserTable)
		var userID int
		if err := tx.QueryRow(query, usr.Username, usr.Password, usr.Email).Scan(&userID); err != nil {
			return err
		}

		query = fmt.Sprintf("INSERT INTO %s (user_id) VALUES ($1)", database.WalletTable)
		if _, err := tx.Exec(query, userID); err != nil {
			return err
		}

		return nil
	})
}

func (s *AuthDB) GetUser(username, password string) (*model.User, error) {
	var user model.User

	if err := s.db.Transaction(func(tx *sqlx.Tx) error {
		query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", database.UserTable)
		return tx.Get(&user, query, username, password)
	}); err != nil {
		return nil, err
	}

	return &user, nil
}
