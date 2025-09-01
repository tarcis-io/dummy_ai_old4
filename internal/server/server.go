// Package server provides a pre-configured HTTP server for the application.
package server

import (
	"net/http"
)

type (
	// Server represents the pre-configured HTTP server for the application.
	Server struct {
		// address specifies the address to listen on.
		address string

		// router specifies the HTTP request multiplexer.
		router *http.ServeMux
	}

	// pageData represents the data that is passed to the HTML template.
	pageData struct {
		// Title specifies the page title.
		Title string

		// WASMPath specifies the path to the WASM file.
		WASMPath string
	}
)

const (
	pageTitleDefault = "DummyAI"
)

func newPageData(wasmPath string) *pageData {
	pageData := &pageData{
		Title:    pageTitleDefault,
		WASMPath: wasmPath,
	}
	return pageData
}
