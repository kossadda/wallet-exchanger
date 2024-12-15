package main

import (
	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/internal/delivery"
	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/internal/repository"
	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/internal/service"
	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/model"
	"github.com/kossadda/wallet-exchanger/share/configs"
	"github.com/kossadda/wallet-exchanger/share/postgres"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func main() {
	cfg, err := configs.NewEnvConfig()
	if err != nil {
		logrus.Fatalf("failed to initialize: %s", err.Error())
	}

	db, err := postgres.NewPostgresDB(cfg)
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	rep := repository.NewRepository(db)
	srv := service.NewService(rep)
	hnd := delivery.NewHandler(srv)

	if err = model.NewServer().Run(cfg.ServerPort, hnd.InitRoutes()); err != nil {
		logrus.Fatalf("error while running http server: %s", err.Error())
	}
}
