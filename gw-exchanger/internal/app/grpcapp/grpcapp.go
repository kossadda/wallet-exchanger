package grpcapp

import (
	"fmt"
	"github.com/kossadda/wallet-exchanger/gw-echanger/internal/storage"
	"log/slog"
	"net"

	"github.com/kossadda/wallet-exchanger/gw-echanger/internal/grpc/exchangegrpc"
	"github.com/kossadda/wallet-exchanger/gw-echanger/internal/service"
	"github.com/kossadda/wallet-exchanger/share/configs"
	"github.com/kossadda/wallet-exchanger/share/database"
	"google.golang.org/grpc"
)

type GRPCApp struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       int
}

func New(log *slog.Logger, port int) *GRPCApp {
	gRPCServer := grpc.NewServer()

	cfg := configs.NewDefaultConfig()
	db, err := database.NewPostgres(cfg)
	if err != nil {
		panic(err)
	}
	strg := storage.New(db)
	srvc := service.New(strg)

	exchangegrpc.Register(gRPCServer, srvc)

	return &GRPCApp{
		log:        log,
		gRPCServer: gRPCServer,
		port:       port,
	}
}

func (a *GRPCApp) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *GRPCApp) Run() error {
	const op = "grpcapp.Run"

	log := a.log.With(slog.String("op", op), slog.Int("port", a.port))

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	log.Info("starting gRPC server", slog.String("address", l.Addr().String()))
	if err := a.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (a *GRPCApp) Stop() {
	const op = "grpcapp.Stop"

	a.log.With(slog.String("op", op)).Info("stopping gRPC server")

	a.gRPCServer.GracefulStop()
}
