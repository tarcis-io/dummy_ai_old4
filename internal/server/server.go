package server

import (
	"bytes"
	"embed"
	"io/fs"
	"log/slog"
	"net/http"
	"text/template"
)

type (
	Server struct {
		address string
		router  *http.ServeMux
		logger  *slog.Logger
	}

	pageData struct {
		Title    string
		WASMPath string
	}
)

const (
	staticDirectory   = "web/static"
	staticPath        = "/static/"
	templatesPattern  = "web/template/*.html"
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

	pageRoutes = map[string]*pageData{
		homePagePath:  homePageData,
		aboutPagePath: aboutPageData,
	}
)

func New(address string, logger *slog.Logger) (*Server, error) {
	staticFS, err := fs.Sub(webFS, staticDirectory)
	if err != nil {
		return nil, err
	}
	staticFileServer := http.FileServer(http.FS(staticFS))
	router := http.NewServeMux()
	router.Handle(staticPath, http.StripPrefix(staticPath, staticFileServer))
	pageTemplates, err := template.ParseFS(webFS, templatesPattern)
	if err != nil {
		return nil, err
	}
	for pagePath, pageData := range pageRoutes {
		var buffer bytes.Buffer
		err = pageTemplates.Execute(&buffer, pageData)
		if err != nil {
			return nil, err
		}
		router.HandleFunc(pagePath, func(w http.ResponseWriter, r *http.Request) {
			w.Write(buffer.Bytes())
		})
	}
	if logger == nil {
		logger = slog.Default()
	}
	server := &Server{
		address: address,
		router:  router,
		logger:  logger,
	}
	return server, nil
}

func (server *Server) Start() error {
	return http.ListenAndServe(server.address, server.router)
}

func newPageData(wasmPath string) *pageData {
	pageData := &pageData{
		Title:    pageTitleDefault,
		WASMPath: wasmPath,
	}
	return pageData
}
