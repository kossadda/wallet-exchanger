package main

import (
	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/pkg/app"
	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/pkg/app/currecywallet"
	"github.com/kossadda/wallet-exchanger/share/configs"
	"github.com/sirupsen/logrus"
)

func main() {
	cfg, err := configs.NewEnvConfig("config.env")
	if err != nil {
		logrus.Fatal(err)
	}

	err = app.Run(currecywallet.New(cfg))
	if err != nil {
		logrus.Fatal(err)
	}
}
