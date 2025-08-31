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
	pageTitleDefault     = "DummyAI"
	homePagePath         = "GET /"
	homePageWASMPath     = "/static/wasm/home.wasm"
	aboutPagePath        = "GET /about"
	aboutPageWASMPath    = "/static/wasm/about.wasm"
)

var (
	//go:embed web/*
	webFS embed.FS

	homePageData  = newPageData(homePageWASMPath)
	aboutPageData = newPageData(aboutPageWASMPath)

	pageRoutes = map[string]*pageData{
		homePagePath:  homePageData,
		aboutPagePath: aboutPageData,
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

func newPageData(wasmPath string) *pageData {
	pageData := &pageData{
		Title:    pageTitleDefault,
		WASMPath: wasmPath,
	}
	return pageData
}
