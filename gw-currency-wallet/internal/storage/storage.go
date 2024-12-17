package storage

import (
	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/internal/storage/auth"
	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/internal/storage/wallet"
	"github.com/kossadda/wallet-exchanger/share/pkg/database"
)

type Storage struct {
	*auth.Auth
	*wallet.Wallet
}

func New(db database.DataBase) *Storage {
	return &Storage{
		Auth:   auth.New(db),
		Wallet: wallet.New(db),
	}
}
