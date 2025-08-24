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

func (server *Server) renderPage(responseWriter http.ResponseWriter, pageData *pageData) error {
	responseWriter.Header().Set("Content-Type", "text/html; chatset=UTF-8")
	return server.pageTemplate.Execute(responseWriter, pageData)
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
