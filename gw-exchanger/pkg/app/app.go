package app

import (
	"log/slog"

	grpcapp "github.com/kossadda/wallet-exchanger/gw-exchanger/internal/app"
	"github.com/kossadda/wallet-exchanger/gw-exchanger/internal/service"
	"github.com/kossadda/wallet-exchanger/gw-exchanger/internal/storage"
	"github.com/kossadda/wallet-exchanger/share/pkg/configs"
	"github.com/kossadda/wallet-exchanger/share/pkg/database"
	"google.golang.org/grpc"
)

type App struct {
	GRPCSrv *grpcapp.GRPCApp
}

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
		GRPCSrv: grpcapp.New(log, gRPCServer, services, appAddr.Port),
	}
}
