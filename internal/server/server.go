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

func (server *Server) renderPage(responseWriter http.ResponseWriter, pageData pageData) error {
	responseWriter.Header().Set("Content-Type", "text/html; charset=UTF-8")
	return server.pageTemplate.Execute(responseWriter, pageData)
}
