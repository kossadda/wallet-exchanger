package service

import (
	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/internal/service/auth"
	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/internal/service/grpcclient"
	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/internal/service/wallet"
	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/internal/storage"
	"github.com/kossadda/wallet-exchanger/share/pkg/configs"
)

type Service struct {
	*auth.Auth
	*wallet.Wallet
	*grpcclient.Exchange
}

func New(repos *storage.Storage, servConfig *configs.ServerConfig) *Service {
	return &Service{
		Auth:     auth.New(repos),
		Wallet:   wallet.New(repos),
		Exchange: grpcclient.New(repos, servConfig),
	}
}
