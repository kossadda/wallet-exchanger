package main

import (
	"log/slog"

	"github.com/kossadda/wallet-exchanger/gw-exchanger/pkg/app"
	"github.com/kossadda/wallet-exchanger/share/pkg/configs"
	"github.com/kossadda/wallet-exchanger/share/pkg/logger"
)

func main() {
	servConf := &configs.ServerConfig{
		Env:         "local",
		TokenExpire: "10h",
		CacheExpire: "1m",
		Servers: map[string]configs.Server{
			"APP": configs.Server{
				Host: "localhost",
				Port: "44044",
			},
		},
	}

	dbConf := &configs.ConfigDB{
		DBHost:     "localhost",
		DBPort:     "5436",
		DBUser:     "postgres",
		DBPassword: "qwerty",
		DBName:     "postgres",
		DBSSLMode:  "disable",
	}

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
