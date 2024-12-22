// Package main contains the entry point for the application.
// It initializes configurations, sets up logging, and runs the WalletApp instance.
package main

import (
	"github.com/gin-gonic/gin"
	"log/slog"

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
	servConf := configs.NewServerEnvConfig("config/local.env")
	dbConf := configs.NewEnvConfigDB("config/database.env")
	log := logger.SetupByEnv(servConf.Env)

	log.Info("start application",
		slog.String("env", servConf.Env),
		slog.Any("server config", servConf),
		slog.Any("postgres config", dbConf),
	)

	application := app.New(log, dbConf, servConf)

	go application.Wallet.MustRun()

	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8282")

	sign := application.Wallet.Stop()
	log.Info("application stopped", slog.String("signal", sign.String()))
}
