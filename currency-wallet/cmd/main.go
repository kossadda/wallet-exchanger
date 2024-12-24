// Package main contains the entry point for the application.
// It initializes configurations, sets up logging, and runs the WalletApp instance.
package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/kossadda/wallet-exchanger/currency-wallet/docs"
	"github.com/kossadda/wallet-exchanger/currency-wallet/pkg/app"
	"github.com/kossadda/wallet-exchanger/share/pkg/configs"
	"github.com/kossadda/wallet-exchanger/share/pkg/logger"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

// @title Currency Wallet API
// @version 1.0
// @description This is the API documentation for the Currency Wallet service.
// @host localhost:8282
// @BasePath /api/v1
func main() {
	servCfg, dbCfg := initConfigs()
	log := logger.SetupByEnv(servCfg.Env)

	log.Info("start application",
		slog.String("env", servCfg.Env),
		slog.Any("server config", servCfg),
		slog.Any("postgres config", dbCfg),
	)

	application := app.New(log, dbCfg, servCfg)

	go application.Wallet.MustRun()

	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8282")

	sign := application.Wallet.Stop()
	log.Info("application stopped", slog.String("signal", sign.String()))
}

func initConfigs() (*configs.ServerConfig, *configs.ConfigDB) {
	fs := flag.NewFlagSet("configs", flag.ContinueOnError)

	servs := fs.String("serv", "", "path to server config file (.env)")
	db := fs.String("db", "", "path to database config file (.env)")

	if err := fs.Parse(os.Args[1:]); err != nil {
		panic(fmt.Sprintf("failed to parse flags: %v", err))
	}

	if *servs == "" {
		panic("server config file path is required")
	}
	if *db == "" {
		panic("database config file path is required")
	}

	if _, err := os.Stat(*servs); os.IsNotExist(err) {
		panic(fmt.Sprintf("server config file does not exist at path: %s", *servs))
	}

	if _, err := os.Stat(*db); os.IsNotExist(err) {
		panic(fmt.Sprintf("database config file does not exist at path: %s", *db))
	}

	return configs.NewServerEnvConfig(*servs), configs.NewEnvConfigDB(*db)
}
