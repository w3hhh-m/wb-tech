package server

import (
	"context"
	"net/http"

	"wb-tech-l0/internal/config"
	"wb-tech-l0/internal/logger"
)

// Server represents the HTTP server
// It is part of the main application
// and doesn't implement any custom interface
// (unlike clients for external services)
type Server struct {
	server *http.Server
	log    logger.Logger
}

// New creates a new HTTP server with the given config, logger and handler
func New(cfg *config.ServerConfig, log logger.Logger, handler http.Handler) *Server {
	return &Server{
		server: &http.Server{
			Addr:         cfg.Address,
			Handler:      handler,
			ReadTimeout:  cfg.ReadTimeout,
			WriteTimeout: cfg.WriteTimeout,
			IdleTimeout:  cfg.IdleTimeout,
		},
		log: log,
	}
}

// Start runs the HTTP server. It returns when the server stops or an error occurs.
// It blocks until error occurs or application is exiting
func (s *Server) Start() error {
	s.log.Debug("Starting HTTP server")
	return s.server.ListenAndServe()
}

// Close gracefully shuts down the HTTP server with the given context
func (s *Server) Close(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
