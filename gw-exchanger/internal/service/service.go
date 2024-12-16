package service

import (
	"context"
	"github.com/kossadda/wallet-exchanger/gw-echanger/internal/storage"

	gen "github.com/kossadda/wallet-exchanger/share/gen/exchange"
)

type Exchanger interface {
	Exchange(ctx context.Context, request *gen.ExchangeRequest) (*gen.ExchangeResponse, error)
}

type Service struct {
	Exchanger
}

func New(strg *storage.Storage) *Service {
	return &Service{
		Exchanger: NewExchangeService(strg),
	}
}
