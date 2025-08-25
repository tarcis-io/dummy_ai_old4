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
	pageDataTitleDefault  = "DummyAI"
	homePageRoutePath     = "GET /"
	homePageDataWASMPath  = "/wasm/home.wasm"
	aboutPageRoutePath    = "GET /about"
	aboutPageDataWASMPath = "/wasm/about.wasm"
)

var (
	//go:embed web/*
	webFS embed.FS

	homePageData  = newPageData(homePageDataWASMPath)
	aboutPageData = newPageData(aboutPageDataWASMPath)

	pageRoutes = map[string]*pageData{
		homePageRoutePath:  homePageData,
		aboutPageRoutePath: aboutPageData,
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
	return server, nil
}

func (server *Server) renderPage(responseWriter http.ResponseWriter, pageData *pageData, statusCode int) error {
	var buffer bytes.Buffer
	err := server.pageTemplate.Execute(&buffer, pageData)
	if err != nil {
		return err
	}
	responseWriter.WriteHeader(statusCode)
	buffer.WriteTo(responseWriter)
	return nil
}

func newPageData(wasmPath string) *pageData {
	pageData := &pageData{
		Title:    pageDataTitleDefault,
		WASMPath: wasmPath,
	}
	return pageData
}
