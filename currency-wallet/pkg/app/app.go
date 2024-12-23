// Package app initializes and manages the core application, including setting up services, handlers, and the server.
// It integrates the wallet service, database connections, and request handling.
package app

import (
	"fmt"
	"log/slog"

	walletapp "github.com/kossadda/wallet-exchanger/currency-wallet/internal/app"
	"github.com/kossadda/wallet-exchanger/currency-wallet/internal/delivery"
	"github.com/kossadda/wallet-exchanger/currency-wallet/internal/service"
	"github.com/kossadda/wallet-exchanger/currency-wallet/internal/storage"
	"github.com/kossadda/wallet-exchanger/share/pkg/configs"
	"github.com/kossadda/wallet-exchanger/share/pkg/database"
)

// App represents the main application structure, containing the WalletApp instance.
type App struct {
	Wallet *walletapp.WalletApp // The WalletApp handles the main application logic.
}

// New creates a new instance of the App, initializes necessary components like the database, services, and handlers.
// It sets up the WalletApp with the given logger, database configuration, and server configuration.
func New(log *slog.Logger, dbConf *configs.ConfigDB, servConf *configs.ServerConfig) *App {
	db, err := database.NewPostgres(dbConf)
	if err != nil {
		panic(fmt.Sprintf("walletApp.New: %v", err))
	}

	services := service.New(storage.New(db), servConf)
	handler := delivery.New(services, log, servConf)

	appAddr, ok := servConf.Servers["APP"]
	if !ok {
		appAddr.Host = "localhost"
		appAddr.Port = configs.DefaultWalletServicePort
	}

	return &App{
		Wallet: walletapp.New(log, handler, appAddr.Port),
	}
}
