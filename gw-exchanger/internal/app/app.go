package app

import (
	"github.com/kossadda/wallet-exchanger/gw-echanger/internal/app/grpcapp"
	"github.com/kossadda/wallet-exchanger/share/configs"
	"log/slog"
)

type App struct {
	GRPCSrv *grpcapp.GRPCApp
}

func New(log *slog.Logger, dbConf *configs.ConfigDB, servConf *configs.ServerConfig) *App {
	grpcApp := grpcapp.New(log, dbConf, servConf)
	return &App{
		GRPCSrv: grpcApp,
	}
}
