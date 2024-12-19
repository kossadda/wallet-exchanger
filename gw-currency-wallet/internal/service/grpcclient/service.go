package grpcclient

import (
	"context"

	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/internal/model"
	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/internal/storage"
	gen "github.com/kossadda/wallet-exchanger/share/gen/exchange"
	"github.com/kossadda/wallet-exchanger/share/pkg/configs"
)

type MainAPI interface {
	GetExchangeRates(ctx context.Context) (*gen.ExchangeRatesResponse, error)
	ExchangeSum(ctx context.Context, input *model.Exchange) ([]float64, error)
	CloseGRPC() error
}

type Exchange struct {
	MainAPI
}

func New(repo *storage.Storage, servConfig *configs.ServerConfig) *Exchange {
	return &Exchange{
		MainAPI: newService(repo, servConfig),
	}
}
