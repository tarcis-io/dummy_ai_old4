package server

import (
	"embed"
	"io/fs"
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

const (
	staticDirectory = "web/static"
	staticPath      = "/static/"
)

var (
	//go:embed web/*
	webFS embed.FS
)

func New(address string, logger *slog.Logger) (*Server, error) {
	staticFS, err := fs.Sub(webFS, staticDirectory)
	if err != nil {
		return nil, err
	}
	staticFileServer := http.FileServer(http.FS(staticFS))
	router := http.NewServeMux()
	router.Handle(staticPath, http.StripPrefix(staticPath, staticFileServer))
	if logger == nil {
		logger = slog.Default()
	}
	server := &Server{
		address: address,
		router:  router,
		logger:  logger,
	}
	return server, nil
}

func (server *Server) Start() error {
	return http.ListenAndServe(server.address, server.router)
}
