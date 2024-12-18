package exchange

import (
	"context"

	gen "github.com/kossadda/wallet-exchanger/share/gen/exchange"
	"github.com/kossadda/wallet-exchanger/share/pkg/database"
)

type MainAPI interface {
	GetExchangeRates(context.Context) (*gen.ExchangeRatesResponse, error)
}

type Exchange struct {
	MainAPI
}

func New(db database.DataBase) *Exchange {
	return &Exchange{
		MainAPI: newStorage(db),
	}
}
