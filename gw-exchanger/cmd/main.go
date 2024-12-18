package main

import (
	"log/slog"

	"github.com/kossadda/wallet-exchanger/gw-exchanger/pkg/app"
	"github.com/kossadda/wallet-exchanger/share/pkg/configs"
	"github.com/kossadda/wallet-exchanger/share/pkg/logger"
)

func main() {
	servCfg := configs.NewServerEnvConfig("config/local.env")
	dbCfg := configs.NewEnvConfigDB("config/database.env")
	log := logger.SetupByEnv(servCfg.Env)

	log.Info("start application",
		slog.String("env", servCfg.Env),
		slog.Any("server config", servCfg),
		slog.Any("postgres config", dbCfg),
	)

	application := app.New(log, dbCfg, servCfg)

	go application.GRPCSrv.MustRun()

	sign := application.GRPCSrv.Stop()
	log.Info("application stopped", slog.String("signal", sign.String()))
}
