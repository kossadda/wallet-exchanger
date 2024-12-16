package main

import (
	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/pkg/app"
	"github.com/kossadda/wallet-exchanger/share/configs"
	"github.com/sirupsen/logrus"
)

func main() {
	cfg := configs.NewEnvConfig("config.env")

	err := app.Run(app.NewCurrecyWallet(cfg))
	if err != nil {
		logrus.Fatal(err)
	}
}
