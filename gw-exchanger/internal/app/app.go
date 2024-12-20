package app

import (
	"fmt"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/kossadda/wallet-exchanger/gw-exchanger/internal/delivery"
	"github.com/kossadda/wallet-exchanger/gw-exchanger/internal/service"
	"github.com/kossadda/wallet-exchanger/gw-exchanger/internal/storage"
	"github.com/kossadda/wallet-exchanger/share/pkg/configs"
	"github.com/kossadda/wallet-exchanger/share/pkg/database"
	"google.golang.org/grpc"
)

type GRPCApp struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       string
}

func New(log *slog.Logger, dbConf *configs.ConfigDB, servConf *configs.ServerConfig) *GRPCApp {
	gRPCServer := grpc.NewServer()

	db, err := database.NewPostgres(dbConf)
	if err != nil {
		panic(err)
	}
	services := service.New(storage.New(db))

	delivery.Register(gRPCServer, services, log)

	appAddr, ok := servConf.Servers["APP"]
	if !ok {
		appAddr.Host = "localhost"
		appAddr.Port = configs.DefaultGrpcPort
	}

	return &GRPCApp{
		log:        log,
		gRPCServer: gRPCServer,
		port:       appAddr.Port,
	}
}

func (a *GRPCApp) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *GRPCApp) Run() error {
	const op = "grpcapp.Run"

	_, err := strconv.Atoi(a.port)
	if err != nil {
		a.port = configs.DefaultGrpcPort
	}
	log := a.log.With(slog.String("op", op), slog.String("port", a.port))

	l, err := net.Listen("tcp", ":"+a.port)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	log.Info("starting gRPC server", slog.String("address", l.Addr().String()))
	if err := a.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (a *GRPCApp) Stop() os.Signal {
	const op = "grpcapp.Stop"

	quitCh := make(chan os.Signal, 1)
	signal.Notify(quitCh, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	sign := <-quitCh

	a.log.With(slog.String("op", op)).Info("stopping gRPC server")

	a.gRPCServer.GracefulStop()

	return sign
}
