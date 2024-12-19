package grpcclient

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/internal/service"
)

type MainAPI interface {
	GetExchangeRates(ctx *gin.Context)
	ExchangeSum(ctx *gin.Context)
	CloseGRPC() error
}

type Exchange struct {
	MainAPI
}

func New(services *service.Service, logger *slog.Logger) *Exchange {
	return &Exchange{
		MainAPI: newHandler(services, logger),
	}
}
