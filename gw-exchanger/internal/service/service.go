// Package service provides high-level services for handling currency exchange operations.
package service

import (
	"github.com/kossadda/wallet-exchanger/gw-exchanger/internal/service/exchange"
	"github.com/kossadda/wallet-exchanger/gw-exchanger/internal/storage"
)

// Service aggregates all services related to currency exchange.
type Service struct {
	*exchange.Exchange
}

// New creates a new instance of Service.
// It initializes the service with the required storage.
func New(strg *storage.Storage) *Service {
	return &Service{
		Exchange: exchange.New(strg),
	}
}
