package app

import (
	"os"
	"os/signal"
	"syscall"

	cw "github.com/kossadda/wallet-exchanger/gw-currency-wallet/internal/app"
	"github.com/kossadda/wallet-exchanger/share/configs"
)

type App interface {
	Start() error
	Stop()
}

func NewCurrecyWallet(cfg *configs.Config) App {
	return cw.New(cfg) 
}

func Run(a App) error {
	var err error
	go func() {
		err = a.Start()
	}()

	quitCh := make(chan os.Signal, 1)
	signal.Notify(quitCh, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	<-quitCh

	a.Stop()

	return err
}
