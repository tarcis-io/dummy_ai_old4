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

func (server *Server) renderPage(responseWriter http.ResponseWriter, pageData pageData) {
	responseWriter.Header().Set("Content-Type", "text/html; charset=UTF-8")
	err := server.pageTemplate.Execute(responseWriter, pageData)
	if err != nil {
		http.Error(responseWriter, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
