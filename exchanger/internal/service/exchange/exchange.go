package exchange

import (
	"context"
	"fmt"
	"strings"

	"github.com/kossadda/wallet-exchanger/exchanger/internal/storage"
	gen "github.com/kossadda/wallet-exchanger/share/gen/exchange"
)

// service handles low-level operations for currency exchange.
type service struct {
	strg *storage.Storage // Database layer
}

// newService creates a new instance of the exchange service.
func newService(strg *storage.Storage) *service {
	return &service{
		strg: strg,
	}
}

// GetExchangeRateForCurrency retrieves the exchange rate between two currencies.
func (es *service) GetExchangeRateForCurrency(ctx context.Context, req *gen.CurrencyRequest) (*gen.ExchangeRateResponse, error) {
	if req.ToCurrency == "" || req.FromCurrency == "" {
		return nil, fmt.Errorf("empty wallet currency request")
	}

	r, err := es.strg.GetExchangeRates(ctx)
	if err != nil {
		return nil, err
	}

	toRates, ok := r.Rates[strings.ToLower(req.ToCurrency)]
	if !ok {
		return nil, fmt.Errorf("invalid output currency %s", req.ToCurrency)
	}
	resRate, ok := toRates.Rate[strings.ToLower(req.FromCurrency)]
	if !ok {
		return nil, fmt.Errorf("invalid input currency %s", req.FromCurrency)
	}

	return &gen.ExchangeRateResponse{
		FromCurrency: req.FromCurrency,
		ToCurrency:   req.ToCurrency,
		Rate:         resRate,
	}, nil
}

// GetExchangeRates retrieves all available exchange rates.
func (es *service) GetExchangeRates(ctx context.Context) (*gen.ExchangeRatesResponse, error) {
	return es.strg.GetExchangeRates(ctx)
}
