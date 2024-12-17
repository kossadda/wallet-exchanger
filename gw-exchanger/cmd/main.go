package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/kossadda/wallet-exchanger/gw-echanger/internal/app"
	"github.com/kossadda/wallet-exchanger/share/configs"
	"github.com/kossadda/wallet-exchanger/share/logger"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	servCfg := configs.NewServerEnvConfig("config/local.env")
	dbCfg := configs.NewEnvConfigDB("config/database.env")

	log := logger.Setup(servCfg.Env)

	log.Info("start application",
		slog.String("env", servCfg.Env),
		slog.Any("cfg", servCfg),
		slog.String("port", servCfg.Port),
	)

	application := app.New(log, dbCfg, servCfg)

	go application.GRPCSrv.MustRun()

	quitCh := make(chan os.Signal, 1)
	signal.Notify(quitCh, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	sign := <-quitCh

	application.GRPCSrv.Stop()
	log.Info("application stopped", slog.String("signal", sign.String()))
}
