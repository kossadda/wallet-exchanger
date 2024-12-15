package app

import (
	"os"
	"os/signal"
	"syscall"

	_ "github.com/lib/pq"
)

type App interface {
	Start() error
	Stop()
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
