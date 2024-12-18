package service

import (
	"github.com/kossadda/wallet-exchanger/gw-echanger/internal/service/exchange"
	"github.com/kossadda/wallet-exchanger/gw-echanger/internal/storage"
)

type Service struct {
	*exchange.Exchange
}

func New(strg *storage.Storage) *Service {
	return &Service{
		Exchange: exchange.New(strg),
	}
}
