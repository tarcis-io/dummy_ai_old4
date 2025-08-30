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
	homePagePath         = "GET /"
	homePageWASMPath     = "/static/wasm/home.wasm"
	aboutPagePath        = "GET /about"
	aboutPageWASMPath    = "/static/wasm/about.wasm"
)

var (
	//go:embed web/*
	webFS embed.FS

	pageRoutes = map[string]*pageData{
		homePagePath:  newPageData(homePageWASMPath),
		aboutPagePath: newPageData(aboutPageWASMPath),
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
	err := http.ListenAndServe(server.address, server.router)
	if err != nil {
		return fmt.Errorf("failed to start server error=%w", err)
	}
	return nil
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
	for pagePath, pageData := range pageRoutes {
		var buffer bytes.Buffer
		err = pageTemplate.Execute(&buffer, pageData)
		if err != nil {
			return fmt.Errorf("failed to execute page template error=%w", err)
		}
		cache := buffer.Bytes()
		server.router.HandleFunc(pagePath, func(responseWriter http.ResponseWriter, request *http.Request) {
			responseWriter.Header().Set("Content-Type", "text/html; charset=UTF-8")
			responseWriter.Write(cache)
		})
	}
	return nil
}

func newPageData(wasmPath string) *pageData {
	pageData := &pageData{
		Title:    pageTitleDefault,
		WASMPath: wasmPath,
	}
	return pageData
}
