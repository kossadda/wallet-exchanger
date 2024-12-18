package app

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/internal/delivery"
	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/internal/service"
	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/internal/storage"
	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/pkg/server"
	"github.com/kossadda/wallet-exchanger/share/pkg/configs"
	"github.com/kossadda/wallet-exchanger/share/pkg/database"
)

type WalletApp struct {
	log    *slog.Logger
	db     database.DataBase
	server *server.Server
	config *configs.ServerConfig
}

func New(log *slog.Logger, dbConf *configs.ConfigDB, servConf *configs.ServerConfig) *WalletApp {
	db, err := database.NewPostgres(dbConf)
	if err != nil {
		panic(err)
	}

	_, err = strconv.Atoi(servConf.Port)
	if err != nil {
		servConf.Port = configs.DefaultWalletServicePort
	}

	services := service.New(storage.New(db))
	handler := delivery.NewHandler(services, log, servConf)

	return &WalletApp{
		log:    log,
		db:     db,
		server: server.New(servConf.Port, handler.InitRoutes()),
		config: servConf,
	}
}

func (a *WalletApp) MustRun() {
	if err := a.Run(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		panic(err)
	}
}

func (a *WalletApp) Run() error {
	const op = "WalletApp.Run"

	log := a.log.With(slog.String("op", op), slog.String("port", a.server.Addr()))

	log.Info("starting CurrencyWallet server", slog.String("address", a.server.Addr()))
	if err := a.server.Run(); err != nil {
		return err
	}

	return nil
}

func (a *WalletApp) Stop() os.Signal {
	const op = "WalletApp.Stop"

	quitCh := make(chan os.Signal, 1)
	signal.Notify(quitCh, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	sign := <-quitCh

	a.log.With(slog.String("op", op)).Info("stopping WalletApp server")

	_ = a.db.Close()
	_ = a.server.Shutdown(context.Background())

	return sign
}
