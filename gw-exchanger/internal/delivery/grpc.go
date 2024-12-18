package delivery

import (
	"context"
	"fmt"
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
	const op = "serverAPI.GetExchangeRateForCurrency"

	log := h.logger.With(
		slog.String("operation", op),
		slog.String("currency on input", req.FromCurrency),
		slog.String("currency on output", req.ToCurrency),
	)

	log.Info("get exchange rate from input currency to output currency")

	rate, err := h.services.GetExchangeRateForCurrency(ctx, req)
	if err != nil {
		h.logger.Error(err.Error())
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return rate, nil
}

func (h *serverAPI) GetExchangeRates(ctx context.Context, req *gen.Empty) (*gen.ExchangeRatesResponse, error) {
	const op = "serverAPI.GetExchangeRates"

	log := h.logger.With(
		slog.String("operation", op),
	)

	log.Info("get all exchange rates")

	rate, err := h.services.GetExchangeRates(ctx)
	if err != nil {
		h.logger.Error(err.Error())
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return rate, nil
}
