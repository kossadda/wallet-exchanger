package grpcclient

import (
	"context"

	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/internal/model"
	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/internal/storage"
	gen "github.com/kossadda/wallet-exchanger/share/gen/exchange"
	"github.com/kossadda/wallet-exchanger/share/pkg/configs"
)

// MainAPI defines the core gRPC client operations: retrieving exchange rates, performing currency exchanges, and closing the gRPC connection.
type MainAPI interface {
	// GetExchangeRates retrieves the exchange rates from the gRPC server or cache.
	GetExchangeRates(ctx context.Context) (*gen.ExchangeRatesResponse, error)

	// ExchangeSum performs a currency exchange for the user.
	ExchangeSum(ctx context.Context, input *model.Exchange) ([]float64, error)

	// CloseGRPC closes the gRPC connection.
	CloseGRPC() error
}

// Exchange represents the gRPC client for performing currency exchange and wallet operations.
type Exchange struct {
	MainAPI
}

// New creates and returns a new instance of Exchange using the provided storage repository and configuration.
func New(repo *storage.Storage, servConfig *configs.ServerConfig) *Exchange {
	return &Exchange{
		MainAPI: newService(repo, servConfig),
	}
}
