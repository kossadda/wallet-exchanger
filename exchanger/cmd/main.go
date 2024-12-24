// Package main provides the entry point for the Wallet Exchanger application.
//
// The application starts a REST API using the Gin framework and exposes the
// Wallet Exchanger API via gRPC.
package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"log/slog"
	"os"

	_ "github.com/kossadda/wallet-exchanger/exchanger/docs"
	"github.com/kossadda/wallet-exchanger/exchanger/pkg/app"
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
	servCfg, dbCfg := initConfigs()
	log := logger.SetupByEnv(servCfg.Env)

	log.Info("start application",
		slog.String("env", servCfg.Env),
		slog.Any("server config", servCfg),
		slog.Any("postgres config", dbCfg),
	)

	application := app.New(log, dbCfg, servCfg)

	go application.GRPCSrv.MustRun()

	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8181")

	sign := application.GRPCSrv.Stop()
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
