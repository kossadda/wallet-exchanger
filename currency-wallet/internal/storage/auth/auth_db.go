// Package auth provides authentication-related functionality, including user creation and retrieval from the database.
package auth

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/kossadda/wallet-exchanger/currency-wallet/internal/model"
	"github.com/kossadda/wallet-exchanger/share/pkg/database"
)

// storage represents the storage layer for authentication, containing a reference to the database.
type storage struct {
	db database.DataBase
}

// newStorage creates and returns a new instance of storage.
func newStorage(db database.DataBase) *storage {
	return &storage{
		db: db,
	}
}

// CreateUser inserts a new user into the database, creating both a user record and an associated wallet.
// It returns an error if the database operation fails.
func (s *storage) CreateUser(usr model.User) error {
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

// GetUser retrieves a user by username and password, returning the user details if found.
// It returns an error if the database operation fails or the user is not found.
func (s *storage) GetUser(username, password string) (*model.User, error) {
	var user model.User

	if err := s.db.Transaction(func(tx *sqlx.Tx) error {
		query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", database.UserTable)
		return tx.Get(&user, query, username, password)
	}); err != nil {
		return nil, err
	}

	return &user, nil
}
