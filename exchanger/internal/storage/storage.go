// Package storage contains the implementation for accessing the application's data layer.
// It integrates different data sources such as the exchange rates and other system configurations.
package storage

import (
	"github.com/kossadda/wallet-exchanger/exchanger/internal/storage/exchange"
	"github.com/kossadda/wallet-exchanger/share/pkg/database"
)

// Storage struct provides access to the application's various data sources.
type Storage struct {
	database.DataBase  // embedded database
	*exchange.Exchange // embedded exchange storage
}

// New creates a new Storage instance with the provided database connection.
func New(db database.DataBase) *Storage {
	return &Storage{
		DataBase: db,
		Exchange: exchange.New(db),
	}
}

// Stop closes the underlying database connection.
func (s *Storage) Stop() {
	_ = s.DataBase.Close()
}
