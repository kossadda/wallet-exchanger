package grpcclient

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/kossadda/wallet-exchanger/currency-wallet/internal/service"
)

// MainAPI is an interface that defines the methods required for handling GRPC client operations.
type MainAPI interface {
	// GetExchangeRates to fetch exchange rates
	GetExchangeRates(ctx *gin.Context)

	// ExchangeSum to exchange currency
	ExchangeSum(ctx *gin.Context)

	// CloseGRPC to close GRPC connection
	CloseGRPC() error
}

// Exchange represents the GRPC client handler with methods for fetching exchange rates and performing exchanges.
type Exchange struct {
	MainAPI
}

// New creates and returns a new instance of the Exchange handler with the provided services and logger.
func New(services *service.Service, logger *slog.Logger) *Exchange {
	return &Exchange{
		MainAPI: newHandler(services, logger),
	}
}
