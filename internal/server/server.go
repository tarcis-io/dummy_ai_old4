// Package server provides a pre-configured HTTP server for the application.
package server

import (
	"net/http"
)

type (
	// Server represents the pre-configured HTTP server for the application.
	Server struct {
		address string
		router  *http.ServeMux
	}

	// pageData represents the data that is passed to the HTML template.
	pageData struct {
		Title    string
		WASMPath string
	}
)
