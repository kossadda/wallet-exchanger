// Package app initializes and runs the main application, handling server startup, shutdown, and application lifecycle management.
package app

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/kossadda/wallet-exchanger/currency-wallet/internal/delivery"
	"github.com/kossadda/wallet-exchanger/currency-wallet/pkg/server"
)

// WalletApp represents the main application for the currency wallet service.
// It contains all components of the application, including logging, database, HTTP server, and route handling.
type WalletApp struct {
	log    *slog.Logger      // Logger instance for logging application events
	hnd    *delivery.Handler // Handler for HTTP routes and middleware
	server *server.Server    // HTTP server for handling incoming requests
}

// New initializes a new WalletApp instance with the provided logger, database connection, route handler, and server config.
// It sets up the application with default values if necessary (e.g., setting default host and port).
func New(log *slog.Logger, hnd *delivery.Handler, port string) *WalletApp {
	return &WalletApp{
		log:    log,
		hnd:    hnd,
		server: server.New(port, hnd.InitRoutes()),
	}
}

// MustRun starts the application and ensures it runs properly, terminating with a panic on non-recoverable errors.
func (a *WalletApp) MustRun() {
	if err := a.Run(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		panic(err)
	}
}

// Run starts the server and logs the start of the application.
// It listens for incoming requests and serves them until an error occurs or the server is stopped.
func (a *WalletApp) Run() error {
	const op = "WalletApp.Run"

	log := a.log.With(slog.String("op", op), slog.String("port", a.server.Addr()))

	log.Info("starting CurrencyWallet server", slog.String("address", a.server.Addr()))
	if err := a.server.Run(); err != nil {
		return err
	}

	return nil
}

// Stop gracefully shuts down the application, waiting for interrupt signals and ensuring resources are released.
// It stops the HTTP server, closes the database connection, and shuts down the gRPC client.
func (a *WalletApp) Stop() os.Signal {
	const op = "WalletApp.Stop"

	quitCh := make(chan os.Signal, 1)
	signal.Notify(quitCh, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	sign := <-quitCh

	a.log.With(slog.String("op", op)).Info("stopping WalletApp server")

	_ = a.server.Shutdown(context.Background())
	a.hnd.Stop()

	return sign
}
