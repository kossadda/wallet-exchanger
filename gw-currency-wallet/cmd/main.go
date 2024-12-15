package main

import (
	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/internal/delivery"
	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/internal/repository"
	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/internal/service"
	"github.com/kossadda/wallet-exchanger/share/configs"
	"github.com/kossadda/wallet-exchanger/share/database"
	"github.com/kossadda/wallet-exchanger/share/server"
	"github.com/sirupsen/logrus"
)

func main() {
	cfg, err := configs.NewEnvConfig()
	if err != nil {
		logrus.Fatalf("failed to initialize: %s", err.Error())
	}

	db, err := database.NewPostgres(cfg)
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	rep := repository.NewRepository(db)
	srv := service.NewService(rep)
	hnd := delivery.NewHandler(srv)

	if err = server.New().Run(cfg.ServerPort, hnd.InitRoutes()); err != nil {
		logrus.Fatalf("error while running http server: %s", err.Error())
	}
}
