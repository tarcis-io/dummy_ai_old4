package server

import (
	"net/http"
)

type (
	Server struct {
		address string
		router  *http.ServeMux
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
	homePageData  = newPageData(homePageWASMPath)
	aboutPageData = newPageData(aboutPageWASMPath)
)

func New(address string) (*Server, error) {
	server := &Server{
		address: address,
		router:  http.NewServeMux(),
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
