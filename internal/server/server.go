package server

import (
	"embed"
	"io/fs"
	"log/slog"
	"net/http"
	"text/template"
)

type (
	Server struct {
		address       string
		router        *http.ServeMux
		pageTemplates *template.Template
		logger        *slog.Logger
	}

	pageData struct {
		Title    string
		WASMPath string
	}
)

const (
	pageTitleDefault  = "DummyAI"
	homePagePath      = "GET /"
	homePageWASMPath  = "/static/wasm/home.wasm"
	aboutPagePath     = "GET /about"
	aboutPageWASMPath = "/static/wasm/about.wasm"
)

var (
	//go:embed web/*
	webFS embed.FS

	homePageData  = newPageData(homePageWASMPath)
	aboutPageData = newPageData(aboutPageWASMPath)
	pageRoutes    = map[string]*pageData{
		homePagePath:  homePageData,
		aboutPagePath: aboutPageData,
	}
)

func New(address string, logger *slog.Logger) (*Server, error) {
	pageTemplates, err := template.ParseFS(webFS, "web/template/*.html")
	if err != nil {
		return nil, err
	}
	if logger == nil {
		logger = slog.Default()
	}
	server := &Server{
		address:       address,
		router:        http.NewServeMux(),
		pageTemplates: pageTemplates,
		logger:        logger,
	}
	staticFS, err := fs.Sub(webFS, "web/static")
	if err != nil {
		return nil, err
	}
	server.router.Handle("/static/", http.StripPrefix("/static/", http.FileServerFS(staticFS)))
	for pagePath, pageData := range pageRoutes {
		server.router.HandleFunc(pagePath, server.makePageHandler(pageData))
	}
	return server, nil
}

func (server *Server) makePageHandler(pageData *pageData) func(http.ResponseWriter, *http.Request) {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		server.renderPage(responseWriter, request, pageData)
	}
}

func (server *Server) renderPage(responseWriter http.ResponseWriter, request *http.Request, pageData *pageData) {
}

func newPageData(wasmPath string) *pageData {
	pageData := &pageData{
		Title:    pageTitleDefault,
		WASMPath: wasmPath,
	}
	return pageData
}
