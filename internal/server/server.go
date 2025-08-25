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
}

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

	pageTemplate = template.Must(template.ParseFS(pageTemplateFS, "web/template/*.html"))
)

func New(address string, logger *slog.Logger) *Server {
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
	return server
}
