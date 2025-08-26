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
	pageTitleDefault     = "DummyAI"
	homePagePath         = "GET /"
	homePageWASMPath     = "/wasm/home.wasm"
	aboutPagePath        = "GET /about"
	aboutPageWASMPath    = "/wasm/about.wasm"
	error404PageWASMPath = "/wasm/error_404.wasm"
	error500PageWASMPath = "/wasm/error_500.wasm"
	rootPath             = "/"
)

var (
	//go:embed web/*
	webFS embed.FS

	homePageData     = newPageData(homePageWASMPath)
	aboutPageData    = newPageData(aboutPageWASMPath)
	error404PageData = newPageData(error404PageWASMPath)
	error500PageData = newPageData(error500PageWASMPath)

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
	server.serveMux.HandleFunc(rootPath, server.error404PageHandler)
}

func (server *Server) makePageHandler(pageData *pageData) func(http.ResponseWriter, *http.Request) {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		server.renderPage(responseWriter, request, pageData, http.StatusOK)
	}
}

func (server *Server) error404PageHandler(responseWriter http.ResponseWriter, request *http.Request) {
	server.renderPage(responseWriter, request, error404PageData, http.StatusNotFound)
}

func (server *Server) renderPage(responseWriter http.ResponseWriter, request *http.Request, pageData *pageData, statusCode int) {
	var buffer bytes.Buffer
	err := server.pageTemplate.Execute(&buffer, pageData)
	if err != nil {
		server.logger.ErrorContext(request.Context(), "failed to render page", "path", request.URL.Path, "error", err)
		if statusCode == http.StatusInternalServerError {
			http.Error(responseWriter, http.StatusText(statusCode), statusCode)
			return
		}
		server.renderPage(responseWriter, request, error500PageData, http.StatusInternalServerError)
		return
	}
	responseWriter.Header().Set("Content-Type", "text/html; charset=UTF-8")
	responseWriter.WriteHeader(statusCode)
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
