package app

import (
	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/configs"
	"os"

	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/internal/delivery"
	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/internal/repository"
	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/internal/service"
	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/model"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func Run() error {
	db, err := model.NewPostgresDB(configs.NewDefaultConfig())
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	rep := repository.NewRepository(db)
	srv := service.NewService(rep)
	hnd := delivery.NewHandler(srv)

	if err = model.NewServer().Run(os.Getenv("PORT"), hnd.InitRoutes()); err != nil {
		logrus.Fatalf("error while running http server: %s", err.Error())
	}

	return err
}
