// Package exchange provides functionality to interact with the exchange database.
// It contains methods for fetching exchange rates from the database.
package exchange

import (
	"context"

	gen "github.com/kossadda/wallet-exchanger/share/gen/exchange"
	"github.com/kossadda/wallet-exchanger/share/pkg/database"
)

// MainAPI defines the main interface for exchange operations.
type MainAPI interface {
	// GetExchangeRates to get exchange rates for all currency
	GetExchangeRates(context.Context) (*gen.ExchangeRatesResponse, error)
}

// Exchange struct implements MainAPI and provides exchange-related services.
type Exchange struct {
	MainAPI
}

// New creates a new Exchange instance with the provided database connection.
func New(db database.DataBase) *Exchange {
	return &Exchange{
		MainAPI: newStorage(db),
	}
}
