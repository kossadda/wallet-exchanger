package app

import (
	"log/slog"

	grpcapp "github.com/kossadda/wallet-exchanger/gw-echanger/internal/app"
	"github.com/kossadda/wallet-exchanger/share/configs"
)

type App struct {
	GRPCSrv *grpcapp.GRPCApp
}

func New(log *slog.Logger, dbConf *configs.ConfigDB, servConf *configs.ServerConfig) *App {
	return &App{
		GRPCSrv: grpcapp.New(log, dbConf, servConf),
	}
}
