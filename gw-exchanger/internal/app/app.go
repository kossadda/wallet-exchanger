package app

import (
	"github.com/kossadda/wallet-exchanger/gw-echanger/internal/app/grpcapp"
	"log/slog"
	"time"
)

type App struct {
	GRPCSrv *grpcapp.GRPCApp
}

func Newsd(log *slog.Logger, grpcPort int, storagePath string, tokenTTL time.Duration) *App {
	grpcApp := grpcapp.New(log, grpcPort)
	return &App{
		GRPCSrv: grpcApp,
	}
}
