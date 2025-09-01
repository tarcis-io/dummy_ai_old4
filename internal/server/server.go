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
	pageTemplatePattern  = "web/template/*.html"
	pageTitleDefault     = "DummyAI"
)

var (
	//go:embed web/*
	webFS embed.FS

	pageRoutes = map[string]*pageData{
		"GET /":      newPageData("/static/wasm/home.wasm"),
		"GET /about": newPageData("/static/wasm/about.wasm"),
	}

	pageHeaders = map[string]string{
		"Content-Type": "text/html; charset=UTF-8",
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
	staticFiles, err := fs.Sub(webFS, staticFilesDirectory)
	if err != nil {
		return fmt.Errorf("failed to open static files directory error=%w", err)
	}
	server.router.Handle(staticFilesPath, http.StripPrefix(staticFilesPath, http.FileServerFS(staticFiles)))
	return nil
}

func (server *Server) registerPageRoutes() error {
	return nil
}

func newPageData(wasmPath string) *pageData {
	pageData := &pageData{
		Title:    pageTitleDefault,
		WASMPath: wasmPath,
	}
	return pageData
}
