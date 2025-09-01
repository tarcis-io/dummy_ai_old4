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
	staticFiles, err := fs.Sub(webFS, staticFilesDirectory)
	if err != nil {
		return fmt.Errorf("failed to open static files directory error=%w", err)
	}
	server.router.Handle(staticFilesPath, http.StripPrefix(staticFilesPath, http.FileServerFS(staticFiles)))
	return nil
}
