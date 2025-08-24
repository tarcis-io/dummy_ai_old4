package server

import (
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

func (server *Server) ListenAndServe() error {
	return http.ListenAndServe(server.address, server.serveMux)
}

func (server *Server) homeHandler(responseWriter http.ResponseWriter, request *http.Request) {
	if request.URL.Path != homeRoute {
		server.logger.ErrorContext(request.Context(), "Route not found", "route", request.URL.Path)
		server.error404Handler(responseWriter, request)
		return
	}
	err := server.renderPage(responseWriter, homePageData)
	if err != nil {
		server.logger.ErrorContext(request.Context(), "Failed to render home page", "error", err)
		server.error500Handler(responseWriter, request)
	}
}

func (server *Server) aboutHandler(responseWriter http.ResponseWriter, request *http.Request) {
	err := server.renderPage(responseWriter, aboutPageData)
	if err != nil {
		server.logger.ErrorContext(request.Context(), "Failed to render about page", "error", err)
		server.error500Handler(responseWriter, request)
	}
}

func (server *Server) error404Handler(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.WriteHeader(http.StatusNotFound)
	err := server.renderPage(responseWriter, error404PageData)
	if err != nil {
		server.logger.ErrorContext(request.Context(), "Failed to render 404 page", "error", err)
		server.error500Handler(responseWriter, request)
	}
}

func (server *Server) error500Handler(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.WriteHeader(http.StatusInternalServerError)
	err := server.renderPage(responseWriter, error500PageData)
	if err != nil {
		server.logger.ErrorContext(request.Context(), "Failed to render 500 page", "error", err)
		http.Error(responseWriter, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func (server *Server) renderPage(responseWriter http.ResponseWriter, pageData *pageData) error {
	responseWriter.Header().Set("Content-Type", "text/html; charset=UTF-8")
	return server.pageTemplate.Execute(responseWriter, pageData)
}

func (server *Server) registerHandlers() {
	server.serveMux.HandleFunc(homeRoute, server.homeHandler)
	server.serveMux.HandleFunc(aboutRoute, server.aboutHandler)
}

const (
	titleDefault     = "DummyAI"
	homeRoute        = "/"
	homeWASMPath     = "/wasm/home.wasm"
	aboutRoute       = "/about"
	aboutWASMPath    = "/wasm/about.wasm"
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

	pageTemplate = template.Must(template.ParseFS(pageTemplateFS, "web/template/*.html"))
)

func New(address string) *Server {
	server := &Server{
		address:      address,
		serveMux:     http.NewServeMux(),
		pageTemplate: pageTemplate,
		logger:       slog.Default(),
	}
	server.registerHandlers()
	return server
}
