// Package exchange implements the core logic for currency exchange services.
package exchange

import (
	"context"

	"github.com/kossadda/wallet-exchanger/gw-exchanger/internal/storage"
	gen "github.com/kossadda/wallet-exchanger/share/gen/exchange"
)

// MainAPI defines the interface for exchange rate operations.
type MainAPI interface {
	GetExchangeRates(context.Context) (*gen.ExchangeRatesResponse, error)
	GetExchangeRateForCurrency(context.Context, *gen.CurrencyRequest) (*gen.ExchangeRateResponse, error)
}

// Exchange is the main implementation of the MainAPI interface.
type Exchange struct {
	MainAPI
}

// New creates a new instance of Exchange.
func New(repo *storage.Storage) *Exchange {
	return &Exchange{
		MainAPI: newService(repo),
	}
}
