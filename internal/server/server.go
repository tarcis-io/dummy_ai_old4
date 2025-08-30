package server

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"
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
	pageTemplatePattern  = "web/template/*.html"
	pageTitleDefault     = "DummyAI"
)

var (
	//go:embed web/*
	webFS embed.FS

	pageRoutes = map[string]string{
		"GET /":      "/static/wasm/home.wasm",
		"GET /about": "/static/wasm/about.wasm",
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
	pageTemplate, err := template.ParseFS(webFS, pageTemplatePattern)
	if err != nil {
		return fmt.Errorf("failed to parse page template error=%w", err)
	}
	for path, wasmPath := range pageRoutes {
		var buffer bytes.Buffer
		err = pageTemplate.Execute(&buffer, struct {
			Title    string
			WASMPath string
		}{
			Title:    pageTitleDefault,
			WASMPath: wasmPath,
		})
		if err != nil {
			return fmt.Errorf("failed to execute page template error=%w", err)
		}
		cache := buffer.Bytes()
		server.router.HandleFunc(path, func(responseWriter http.ResponseWriter, request *http.Request) {
			responseWriter.Header().Set("Content-Type", "text/html")
			responseWriter.Write(cache)
		})
	}
	return nil
}
