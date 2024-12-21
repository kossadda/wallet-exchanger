package app

import (
	"fmt"
	"log/slog"

	walletapp "github.com/kossadda/wallet-exchanger/gw-currency-wallet/internal/app"
	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/internal/delivery"
	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/internal/service"
	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/internal/storage"
	"github.com/kossadda/wallet-exchanger/share/pkg/configs"
	"github.com/kossadda/wallet-exchanger/share/pkg/database"
)

type App struct {
	Wallet *walletapp.WalletApp
}

func New(log *slog.Logger, dbConf *configs.ConfigDB, servConf *configs.ServerConfig) *App {
	db, err := database.NewPostgres(dbConf)
	if err != nil {
		panic(fmt.Sprintf("walletApp.New: %v", err))
	}

	services := service.New(storage.New(db), servConf)
	handler := delivery.New(services, log, servConf)

	return &App{
		Wallet: walletapp.New(log, db, handler, servConf),
	}
}
