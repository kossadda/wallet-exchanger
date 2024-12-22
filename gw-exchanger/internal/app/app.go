// Package app provides the main application logic for running and managing a gRPC server.
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
	"github.com/kossadda/wallet-exchanger/share/pkg/configs"
	"google.golang.org/grpc"
)

// GRPCApp represents the gRPC application, responsible for managing
// the lifecycle of a gRPC server including startup, shutdown, and logging.
type GRPCApp struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       string
}

// New initializes a new instance of GRPCApp with the provided dependencies.
// It also registers the service delivery layer with the gRPC server.
func New(log *slog.Logger, gRPCServer *grpc.Server, services *service.Service, port string) *GRPCApp {
	delivery.Register(gRPCServer, services, log)

	return &GRPCApp{
		log:        log,
		gRPCServer: gRPCServer,
		port:       port,
	}
}

// MustRun starts the gRPC server and panics if an error occurs.
// It acts as a wrapper for the Run method.
func (a *GRPCApp) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

// Run starts the gRPC server on the configured port. If the port is invalid,
// it defaults to the application configuration's default gRPC port.
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

// Stop gracefully stops the gRPC server in response to system signals
// such as SIGINT, SIGTERM, or SIGHUP.
func (a *GRPCApp) Stop() os.Signal {
	const op = "grpcapp.Stop"

	quitCh := make(chan os.Signal, 1)
	signal.Notify(quitCh, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	sign := <-quitCh

	a.log.With(slog.String("op", op)).Info("stopping gRPC server")

	a.gRPCServer.GracefulStop()

	return sign
}
