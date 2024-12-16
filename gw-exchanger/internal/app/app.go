package app

import (
	"log/slog"
	"time"

	"github.com/kossadda/wallet-exchanger/gw-echanger/internal/app/grpcapp"
)

type App struct {
	GRPCSrv *grpcapp.GRPCApp
}

func New(log *slog.Logger, grpcPort int, storagePath string, tokenTTL time.Duration) *App {
	grpcApp := grpcapp.New(log, grpcPort)
	return &App{
		GRPCSrv: grpcApp,
	}
}
