package server

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
)

const (
	staticFilesDirectory = "web/static"
	staticFilesPath      = "/static/"
)

var (
	//go:embed web/*
	webFS embed.FS
)

type (
	Server struct {
		address string
		router  *http.ServeMux
	}
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
	return server, nil
}

func (server *Server) Start() error {
	err := http.ListenAndServe(server.address, server.router)
	if err != nil {
		return fmt.Errorf("failed to start server address=%s error=%w", server.address, err)
	}
	return nil
}

func (server *Server) registerStaticFiles() error {
	staticFilesFS, err := fs.Sub(webFS, staticFilesDirectory)
	if err != nil {
		return fmt.Errorf("failed to register static files error=%w", err)
	}
	server.router.Handle(staticFilesPath, http.StripPrefix(staticFilesPath, http.FileServerFS(staticFilesFS)))
	return nil
}
