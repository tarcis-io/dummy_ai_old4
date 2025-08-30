package server

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
)

type (
	Server struct {
		address string
		router  *http.ServeMux
	}
)

const (
	staticFilesDirectory = "web/static"
	staticFilesPath      = "/static/"
)

var (
	//go:embed web/*
	webFS embed.FS
)

func New(address string) (*Server, error) {
	server := &Server{
		address: address,
		router:  http.NewServeMux(),
	}
	err := server.registerStaticFiles()
	if err != nil {
		return nil, err
	}
	err = server.registerRoutes()
	if err != nil {
		return nil, fmt.Errorf("failed to register routes error=%w", err)
	}
	return server, nil
}

func (server *Server) Start() error {
	return http.ListenAndServe(server.address, server.router)
}

func (server *Server) registerStaticFiles() error {
	staticFiles, err := fs.Sub(webFS, staticFilesDirectory)
	if err != nil {
		return fmt.Errorf("failed to register static files error=%w", err)
	}
	server.router.Handle(staticFilesPath, http.StripPrefix(staticFilesPath, http.FileServerFS(staticFiles)))
	return nil
}

func (server *Server) registerRoutes() error {
	return nil
}
