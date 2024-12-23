// Package service contains business logic and data operations, interacting with external systems such as databases and APIs.
package service

import (
	"github.com/kossadda/wallet-exchanger/currency-wallet/internal/service/auth"
	"github.com/kossadda/wallet-exchanger/currency-wallet/internal/service/grpcclient"
	"github.com/kossadda/wallet-exchanger/currency-wallet/internal/service/wallet"
	"github.com/kossadda/wallet-exchanger/currency-wallet/internal/storage"
	"github.com/kossadda/wallet-exchanger/share/pkg/configs"
)

// Service bundles authentication, wallet, and exchange services together.
type Service struct {
	*storage.Storage
	*auth.Auth           // Embeds Auth for authentication-related operations
	*wallet.Wallet       // Embeds Wallet for wallet-related operations
	*grpcclient.Exchange // Embeds Exchange to get exchange-rates data
}

// New creates and returns a new instance of Service, combining Auth, Wallet, and Exchange services.
func New(repos *storage.Storage, servConfig *configs.ServerConfig) *Service {
	return &Service{
		Storage:  repos,
		Auth:     auth.New(repos),
		Wallet:   wallet.New(repos),
		Exchange: grpcclient.New(repos, servConfig),
	}
}

// Stop shuts down the underlying storage and releases associated resources.
func (s *Service) Stop() {
	s.Storage.Stop()
}
