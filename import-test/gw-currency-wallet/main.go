package main

import (
	"github.com/kossadda/wallet-exchanger/currency-wallet/pkg/app"
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
				Port: "8080",
			},
			"GRPC": configs.Server{
				Host: "localhost",
				Port: "44044",
			},
			"CACHE": configs.Server{
				Host: "localhost",
				Port: "6379",
			},
		},
	}

	log := logger.SetupByEnv(servConf.Env)
	dbConf := &configs.ConfigDB{
		DBHost:     "localhost",
		DBPort:     "5436",
		DBUser:     "postgres",
		DBPassword: "qwerty",
		DBName:     "postgres",
		DBSSLMode:  "disable",
	}

	application := app.New(log, dbConf, servConf)
	go application.Wallet.MustRun()

	_ = application.Wallet.Stop()
}
