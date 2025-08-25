package server

import (
	"embed"
	"net/http"
	"text/template"
)

type (
	Server struct {
		address      string
		serveMux     *http.ServeMux
		pageTemplate *template.Template
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
	homePageData  = newPageData(homePageDataWASMPath)
	aboutPageData = newPageData(aboutPageDataWASMPath)

	pageRoutes = map[string]*pageData{
		homePageRoutePath:  homePageData,
		aboutPageRoutePath: aboutPageData,
	}

	//go:embed web/*
	webFS embed.FS
)

func New(address string) (*Server, error) {
	pageTemplate, err := template.ParseFS(webFS, "web/template/*.html")
	if err != nil {
		return nil, err
	}
	server := &Server{
		address:      address,
		serveMux:     http.NewServeMux(),
		pageTemplate: pageTemplate,
	}
	return server, nil
}

func newPageData(wasmPath string) *pageData {
	pageData := &pageData{
		Title:    pageDataTitleDefault,
		WASMPath: wasmPath,
	}
	return pageData
}
