package main

import (
	"log/slog"

	"github.com/kossadda/wallet-exchanger/gw-exchanger/pkg/app"
	"github.com/kossadda/wallet-exchanger/share/pkg/configs"
	"github.com/kossadda/wallet-exchanger/share/pkg/logger"
)

func main() {
	servConf := configs.NewServerEnvConfig("config/local.env")
	dbConf := configs.NewEnvConfigDB("config/database.env")
	log := logger.SetupByEnv(servConf.Env)

	log.Info("start application",
		slog.String("env", servConf.Env),
		slog.Any("server config", servConf),
		slog.Any("postgres config", dbConf),
	)

	application := app.New(log, dbConf, servConf)

	go application.GRPCSrv.MustRun()

	sign := application.GRPCSrv.Stop()
	log.Info("application stopped", slog.String("signal", sign.String()))
}
