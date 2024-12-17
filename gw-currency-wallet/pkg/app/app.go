package app

import (
	"log/slog"

	walletapp "github.com/kossadda/wallet-exchanger/gw-currency-wallet/internal/app"
	"github.com/kossadda/wallet-exchanger/share/configs"
)

type App struct {
	Wallet *walletapp.WalletApp
}

func New(log *slog.Logger, dbConf *configs.ConfigDB, servConf *configs.ServerConfig) *App {
	return &App{
		Wallet: walletapp.New(log, dbConf, servConf),
	}
}
