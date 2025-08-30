package server

import (
	"log/slog"
	"net/http"
)

type (
	Server struct {
		address string
		router  *http.ServeMux
		logger  *slog.Logger
	}
)

func New(address string, logger *slog.Logger) (*Server, error) {
	if logger == nil {
		logger = slog.Default()
	}
	server := &Server{
		address: address,
		logger:  logger,
	}
	return server, nil
}

func (server *Server) Start() error {
	return http.ListenAndServe(server.address, server.router)
}
