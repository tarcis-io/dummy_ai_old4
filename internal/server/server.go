package server

import (
	"net/http"
	"text/template"
)

type (
	Server struct {
		address      string
		handler      *http.ServeMux
		pageTemplate *template.Template
	}

	pageData struct {
		Title    string
		WASMPath string
	}
)

func (server *Server) ListenAndServe() error {
	return http.ListenAndServe(server.address, server.handler)
}

func (server *Server) homeHandler(responseWriter http.ResponseWriter, request *http.Request) {
	pageData := &pageData{
		Title:    "DummyAI",
		WASMPath: "/wasm/home.wasm",
	}
	err := server.renderPage(responseWriter, pageData)
	if err != nil {
		server.error500Handler(responseWriter, request)
	}
}

func (server *Server) error500Handler(responseWriter http.ResponseWriter, _ *http.Request) {
	responseWriter.WriteHeader(http.StatusInternalServerError)
	pageData := &pageData{
		Title:    "DummyAI",
		WASMPath: "/wasm/error_500.wasm",
	}
	err := server.renderPage(responseWriter, pageData)
	if err != nil {
		http.Error(responseWriter, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func (server *Server) renderPage(responseWriter http.ResponseWriter, pageData *pageData) error {
	responseWriter.Header().Set("Content-Type", "text/html; charset=UTF-8")
	return server.pageTemplate.Execute(responseWriter, pageData)
}
