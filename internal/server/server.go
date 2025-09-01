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

	pageData struct {
		Title    string
		WASMPath string
	}
)

const (
	staticFilesDirectory = "web/static"
	staticFilesPath      = "/static/"
)

var (
	//go:embed web/*
	webFS embed.FS

	pageHeaders = map[string]string{}
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
	err = server.registerPageRoutes()
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
		return fmt.Errorf("failed to open static files directory directory=%s error=%w", staticFilesDirectory, err)
	}
	server.router.Handle(staticFilesPath, http.StripPrefix(staticFilesPath, http.FileServerFS(staticFilesFS)))
	return nil
}

func (server *Server) registerPageRoutes() error {
	return nil
}

func newPageHandler(pageCache []byte) func(http.ResponseWriter, *http.Request) {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		for header, value := range pageHeaders {
			responseWriter.Header().Set(header, value)
		}
		responseWriter.Write(pageCache)
	}
}
