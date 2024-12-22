// Package main provides the entry point for the Wallet Exchanger application.
//
// The application starts a REST API using the Gin framework and exposes the
// Wallet Exchanger API via gRPC.
package main

import (
	"github.com/gin-gonic/gin"
	"log/slog"

	_ "github.com/kossadda/wallet-exchanger/gw-exchanger/docs"
	"github.com/kossadda/wallet-exchanger/gw-exchanger/pkg/app"
	"github.com/kossadda/wallet-exchanger/share/pkg/configs"
	"github.com/kossadda/wallet-exchanger/share/pkg/logger"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

// @title Wallet Exchanger API
// @version 1.0
// @description This is the API documentation for the Wallet Exchanger service.
// @host localhost:8181
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

	go application.GRPCSrv.MustRun()

	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8181")

	sign := application.GRPCSrv.Stop()
	log.Info("application stopped", slog.String("signal", sign.String()))
}
