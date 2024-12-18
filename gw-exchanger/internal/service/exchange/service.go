package exchange

import (
	"context"

	"github.com/kossadda/wallet-exchanger/gw-exchanger/internal/storage"
	gen "github.com/kossadda/wallet-exchanger/share/gen/exchange"
)

type MainAPI interface {
	GetExchangeRates(context.Context) (*gen.ExchangeRatesResponse, error)
	GetExchangeRateForCurrency(context.Context, *gen.CurrencyRequest) (*gen.ExchangeRateResponse, error)
}

type Exchange struct {
	MainAPI
}

func New(repo *storage.Storage) *Exchange {
	return &Exchange{
		MainAPI: newService(repo),
	}
}
