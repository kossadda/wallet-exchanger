package service

import (
	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/internal/service/auth"
	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/internal/service/wallet"
	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/internal/storage"
)

type Service struct {
	*auth.Auth
	*wallet.Wallet
}

func New(repos *storage.Storage) *Service {
	return &Service{
		Auth:   auth.New(repos),
		Wallet: wallet.New(repos),
	}
}
