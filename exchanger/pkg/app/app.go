// Package app provides the core application setup, integrating all the essential components
// such as the gRPC server, database connection, and service initialization.
package app

import (
	"log/slog"

	grpcapp "github.com/kossadda/wallet-exchanger/exchanger/internal/app"
	"github.com/kossadda/wallet-exchanger/exchanger/internal/service"
	"github.com/kossadda/wallet-exchanger/exchanger/internal/storage"
	"github.com/kossadda/wallet-exchanger/share/pkg/configs"
	"github.com/kossadda/wallet-exchanger/share/pkg/database"
	"google.golang.org/grpc"
)

// App represents the main application structure, which includes all necessary components
// such as the gRPC server application.
type App struct {
	GRPCSrv *grpcapp.GRPCApp
}

// New initializes and returns a new instance of the App, setting up the gRPC server,
// database connection, and services.
func New(log *slog.Logger, dbConf *configs.ConfigDB, servConf *configs.ServerConfig) *App {
	gRPCServer := grpc.NewServer()

	db, err := database.NewPostgres(dbConf)
	if err != nil {
		panic(err)
	}
	services := service.New(storage.New(db))

	appAddr, ok := servConf.Servers["APP"]
	if !ok {
		appAddr.Host = "localhost"
		appAddr.Port = configs.DefaultGrpcPort
	}

	return &App{
		GRPCSrv: grpcapp.New(gRPCServer, log, services, appAddr.Port),
	}
}
