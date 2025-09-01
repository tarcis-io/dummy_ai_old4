// Package server provides a pre-configured HTTP server for the application.
package server

import (
	"embed"
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
	// staticFilesDirectory is the directory containing the static files.
	staticFilesDirectory = "web/static"

	// staticFilesPathPrefix is the path prefix for the static files.
	staticFilesPathPrefix = "/static/"

	// pageTitleDefault is the default value for the page title.
	pageTitleDefault = "DummyAI"
)

var (
	// webFS is the embedded file system for the web directory.
	//go:embed web
	webFS embed.FS
)

// newPageData creates and returns a new pageData instance.
func newPageData(wasmPath string) *pageData {
	pageData := &pageData{
		Title:    pageTitleDefault,
		WASMPath: wasmPath,
	}
	return pageData
}
