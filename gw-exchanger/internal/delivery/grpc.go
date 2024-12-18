package delivery

import (
	"context"
	"log/slog"

	"github.com/kossadda/wallet-exchanger/gw-exchanger/internal/service"
	gen "github.com/kossadda/wallet-exchanger/share/gen/exchange"
	"google.golang.org/grpc"
)

type serverAPI struct {
	gen.UnimplementedExchangeServiceServer
	services *service.Service
	logger   *slog.Logger
}

func Register(gRPC *grpc.Server, service *service.Service, log *slog.Logger) {
	gen.RegisterExchangeServiceServer(gRPC, &serverAPI{
		services: service,
		logger:   log,
	})
}

func (h *serverAPI) GetExchangeRateForCurrency(ctx context.Context, req *gen.CurrencyRequest) (*gen.ExchangeRateResponse, error) {
	rate, err := h.services.GetExchangeRateForCurrency(ctx, req)
	if err != nil {
		h.logger.Error(err.Error())
	}

	return rate, err
}

func (h *serverAPI) GetExchangeRates(ctx context.Context, req *gen.Empty) (*gen.ExchangeRatesResponse, error) {
	rate, err := h.services.GetExchangeRates(ctx)
	if err != nil {
		h.logger.Error(err.Error())
	}

	return rate, err
}
