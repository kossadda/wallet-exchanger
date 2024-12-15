package main

import (
	"fmt"
	"os"

	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/pkg/app"
	"github.com/kossadda/wallet-exchanger/gw-currency-wallet/pkg/app/currecywallet"
	"github.com/kossadda/wallet-exchanger/share/configs"
)

func main() {
	cfg := configs.NewDefaultConfig()
	err := app.Run(currecywallet.New(cfg))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
