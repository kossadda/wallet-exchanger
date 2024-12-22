// Package delivery provides the gRPC server implementation for the Wallet Exchanger service.
// This includes methods for querying exchange rates between currencies and for retrieving all exchange rates.
package delivery

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/kossadda/wallet-exchanger/exchanger/internal/service"
	gen "github.com/kossadda/wallet-exchanger/share/gen/exchange"
	"google.golang.org/grpc"
)

// serverAPI implements the gRPC server interface for the exchange service.
type serverAPI struct {
	gen.UnimplementedExchangeServiceServer
	services *service.Service
	logger   *slog.Logger
}

// Register sets up the gRPC server to handle exchange-related requests.
func Register(gRPC *grpc.Server, service *service.Service, log *slog.Logger) {
	gen.RegisterExchangeServiceServer(gRPC, &serverAPI{
		services: service,
		logger:   log,
	})
}

// GetExchangeRateForCurrency fetches exchange rate for input currency from the database for output currency and returns them.
// @Summary Get exchange rate between two currencies
// @Description Retrieve the exchange rate from one currency to another
// @Tags Exchange
// @Accept json
// @Produce json
// @Param from_currency query string true "Currency to convert from"
// @Param to_currency query string true "Currency to convert to"
// @Success 200
// @Failure 400
// @Router /api/v1/exchange-rate [get]
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

// GetExchangeRates fetches exchange rates for from the database and returns them.
// @Summary Get all available exchange rates
// @Description Retrieve all exchange rates
// @Tags Exchange
// @Accept json
// @Produce json
// @Success 200
// @Failure 500
// @Router /api/v1/exchange-rates [get]
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
