package service

import (
	"context"

	"github.com/kossadda/wallet-exchanger/gw-echanger/internal/storage"
	gen "github.com/kossadda/wallet-exchanger/share/gen/exchange"
)

type ExchangeService struct {
	strg *storage.Storage
}

func (e *ExchangeService) Exchange(ctx context.Context, request *gen.ExchangeRequest) (*gen.ExchangeResponse, error) {
	return e.strg.Exchange(ctx, request)
}

func NewExchangeService(strg *storage.Storage) *ExchangeService {
	return &ExchangeService{
		strg: strg,
	}
}
