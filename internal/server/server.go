package server

import (
	"bytes"
	"embed"
	"log/slog"
	"net/http"
	"text/template"
)

type (
	Server struct {
		address      string
		serveMux     *http.ServeMux
		pageTemplate *template.Template
		logger       *slog.Logger
	}

	pageData struct {
		Title    string
		WASMPath string
	}
)

const (
	pageTitleDefault  = "DummyAI"
	homePagePath      = "GET /"
	homePageWASMPath  = "/wasm/home.wasm"
	aboutPagePath     = "GET /about"
	aboutPageWASMPath = "/wasm/about.wasm"
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

func New(address string, logger *slog.Logger) (*Server, error) {
	pageTemplate, err := template.ParseFS(webFS, "web/template/*.html")
	if err != nil {
		return nil, err
	}
	if logger == nil {
		logger = slog.Default()
	}
	server := &Server{
		address:      address,
		serveMux:     http.NewServeMux(),
		pageTemplate: pageTemplate,
		logger:       logger,
	}
	server.registerHandlers()
	return server, nil
}

func (server *Server) registerHandlers() {
	for path, pageData := range pageRoutes {
		server.serveMux.HandleFunc(path, server.makePageHandler(pageData))
	}
}

func (server *Server) makePageHandler(pageData *pageData) func(http.ResponseWriter, *http.Request) {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		server.renderPage(responseWriter, request, pageData)
	}
}

func (server *Server) renderPage(responseWriter http.ResponseWriter, request *http.Request, pageData *pageData) {
	var buffer bytes.Buffer
	err := server.pageTemplate.Execute(&buffer, pageData)
	if err != nil {
		server.logger.ErrorContext(request.Context(), "failed to render page", "path", request.URL.Path, "error", err)
		http.Error(responseWriter, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	responseWriter.Header().Set("Content-Type", "text/html; charset=UTF-8")
	_, err = buffer.WriteTo(responseWriter)
	if err != nil {
		server.logger.WarnContext(request.Context(), "failed to write response", "path", request.URL.Path, "error", err)
	}
}

func newPageData(wasmPath string) *pageData {
	pageData := &pageData{
		Title:    pageTitleDefault,
		WASMPath: wasmPath,
	}
	return pageData
}
