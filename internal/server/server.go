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
	titleDefault     = "DummyAI"
	homePath         = "GET /"
	homeWASMPath     = "/wasm/home.wasm"
	aboutPath        = "GET /about"
	aboutWASMPath    = "/wasm/about.wasm"
	catchAllPath     = "/"
	error404WASMPath = "/wasm/error_404.wasm"
	error500WASMPath = "/wasm/error_500.wasm"
)

var (
	homePageData = &pageData{
		Title:    titleDefault,
		WASMPath: homeWASMPath,
	}

	aboutPageData = &pageData{
		Title:    titleDefault,
		WASMPath: aboutWASMPath,
	}

	error404PageData = &pageData{
		Title:    titleDefault,
		WASMPath: error404WASMPath,
	}

	error500PageData = &pageData{
		Title:    titleDefault,
		WASMPath: error500WASMPath,
	}

	//go:embed web/template/*.html
	pageTemplateFS embed.FS
)

func New(address string, logger *slog.Logger) (*Server, error) {
	pageTemplate, err := template.ParseFS(pageTemplateFS, "web/template/*.html")
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

func (server *Server) createPageHandler(pageData *pageData) func(http.ResponseWriter, *http.Request) {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		err := server.renderPage(responseWriter, pageData, http.StatusOK)
		if err != nil {
			server.logger.ErrorContext(request.Context(), "failed to render page", "path", request.URL.Path, "error", err)
			server.error500Handler(responseWriter, request)
		}
	}
}

func (server *Server) catchAllHandler(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		server.logger.ErrorContext(request.Context(), "method not allowed", "path", request.URL.Path, "method", request.Method)
		http.Error(responseWriter, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	server.error404Handler(responseWriter, request)
}

func (server *Server) error404Handler(responseWriter http.ResponseWriter, request *http.Request) {
	err := server.renderPage(responseWriter, error404PageData, http.StatusNotFound)
	if err != nil {
		server.logger.ErrorContext(request.Context(), "failed to render error 404 page", "path", request.URL.Path, "error", err)
		server.error500Handler(responseWriter, request)
	}
}

func (server *Server) error500Handler(responseWriter http.ResponseWriter, request *http.Request) {
	err := server.renderPage(responseWriter, error500PageData, http.StatusInternalServerError)
	if err != nil {
		server.logger.ErrorContext(request.Context(), "failed to render error 500 page", "path", request.URL.Path, "error", err)
		http.Error(responseWriter, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func (server *Server) renderPage(responseWriter http.ResponseWriter, pageData *pageData, statusCode int) error {
	var buffer bytes.Buffer
	err := server.pageTemplate.Execute(&buffer, pageData)
	if err != nil {
		return err
	}
	responseWriter.Header().Set("Content-Type", "text/html; charset=UTF-8")
	responseWriter.WriteHeader(statusCode)
	buffer.WriteTo(responseWriter)
	return nil
}

func (server *Server) registerHandlers() {
	server.serveMux.HandleFunc(homePath, server.createPageHandler(homePageData))
	server.serveMux.HandleFunc(aboutPath, server.createPageHandler(aboutPageData))
	server.serveMux.HandleFunc(catchAllPath, server.catchAllHandler)
}
