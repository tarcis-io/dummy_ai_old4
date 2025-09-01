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
	staticFilesDirectory       = "web/static"
	staticFilesPath            = "/static/"
	pageTemplatePattern        = "web/template/*.html"
	pageHeaderContentType      = "Content-Type"
	pageHeaderContentTypeValue = "text/html; charset=UTF-8"
	pageTitleDefault           = "DummyAI"
	homePagePath               = "GET /"
	homePageWASMPath           = "/static/wasm/home.wasm"
	aboutPagePath              = "GET /about"
	aboutPageWASMPath          = "/static/wasm/about.wasm"
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
		return fmt.Errorf("failed to register static files error=%w", err)
	}
	server.router.Handle(staticFilesPath, http.StripPrefix(staticFilesPath, http.FileServerFS(staticFilesFS)))
	return nil
}

func (server *Server) registerPageRoutes() error {
	pageTemplate, err := template.ParseFS(webFS, pageTemplatePattern)
	if err != nil {
		return fmt.Errorf("failed to parse page templates pattern=%s error=%w", pageTemplatePattern, err)
	}
	pageRoutes := map[string]*pageData{
		homePagePath:  newPageData(homePageWASMPath),
		aboutPagePath: newPageData(aboutPageWASMPath),
	}
	pageHeaders := map[string]string{
		pageHeaderContentType: pageHeaderContentTypeValue,
	}
	for pagePath, pageData := range pageRoutes {
		var buffer bytes.Buffer
		err := pageTemplate.Execute(&buffer, pageData)
		if err != nil {
			return fmt.Errorf("failed to execute page template path=%s wasmPath=%s error=%w", pagePath, pageData.WASMPath, err)
		}
		pageCache := buffer.Bytes()
		server.router.HandleFunc(pagePath, func(responseWriter http.ResponseWriter, request *http.Request) {
			for pageHeader, pageHeaderValue := range pageHeaders {
				responseWriter.Header().Set(pageHeader, pageHeaderValue)
			}
			responseWriter.Write(pageCache)
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
