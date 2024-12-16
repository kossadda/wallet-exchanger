package app

import (
	"context"

	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/internal/delivery"
	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/internal/service"
	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/internal/storage"
	"github.com/kossadda/wallet-exchanger/share/configs"
	"github.com/kossadda/wallet-exchanger/share/database"
	"github.com/kossadda/wallet-exchanger/share/server"
	"github.com/sirupsen/logrus"
)

type CurrecyWallet struct {
	cfg *configs.Config
	db  database.DataBase
	srv *server.Server
}

func New(cfg *configs.Config) *CurrecyWallet {
	db, err := database.NewPostgres(cfg)
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	srv := server.New()

	return &CurrecyWallet{
		cfg: cfg,
		db:  db,
		srv: srv,
	}
}

func (cw *CurrecyWallet) Start() error {
	rep := storage.NewRepository(cw.db)
	srv := service.NewService(rep)
	hnd := delivery.NewHandler(srv)

	if err := cw.srv.Run(cw.cfg.ServerPort, hnd.InitRoutes()); err != nil {
		return err
	}

	return nil
}

func (cw *CurrecyWallet) Stop() {
	cw.db.Close()
	cw.srv.Shutdown(context.Background())
}
