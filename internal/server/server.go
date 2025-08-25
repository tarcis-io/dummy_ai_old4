package server

import (
	"net/http"
)

type (
	Server struct {
		address  string
		serveMux *http.ServeMux
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
)

func New(address string) (*Server, error) {
	server := &Server{
		address:  address,
		serveMux: http.NewServeMux(),
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
