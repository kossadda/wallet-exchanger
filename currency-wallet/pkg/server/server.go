// Package server provides an abstraction for running an HTTP server.
// It manages the server configuration, including timeouts and shutdown.
package server

import (
	"context"
	"net/http"
	"time"
)

// Server represents an HTTP server with configurable timeouts and handler.
// It encapsulates an *http.Server instance.
type Server struct {
	httpServer *http.Server
}

// New creates a new Server instance with the specified port and HTTP handler.
// It initializes the server with default timeout settings and a handler.
func New(port string, handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:           ":" + port,
			Handler:        handler,
			MaxHeaderBytes: 1 << 20,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
		},
	}
}

// Addr returns the address (including port) that the server is listening on.
func (s *Server) Addr() string {
	return s.httpServer.Addr
}

// Run starts the HTTP server and listens for incoming requests.
// It blocks until the server shuts down.
func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

// Shutdown gracefully shuts down the HTTP server, allowing existing requests to complete.
// It takes a context to control the shutdown timeout.
func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
