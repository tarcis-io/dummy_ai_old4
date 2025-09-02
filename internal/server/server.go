// Package server provides a pre-configured HTTP server for the application.
package server

import (
	"embed"
	"fmt"
	"io/fs"
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

	// pageTemplatePattern is the directory pattern for the HTML page templates.
	pageTemplatePattern = "web/template/*.html"

	// pageTitleDefault is the default value for the page title.
	pageTitleDefault = "DummyAI"

	// homePagePath is the path for the home page.
	homePagePath = "GET /"

	// homePageWASMPath is the path for the home page WASM file.
	homePageWASMPath = "/static/wasm/home.wasm"

	// aboutPagePath is the path for the about page.
	aboutPagePath = "GET /about"

	// aboutPageWASMPath is the path for the about page WASM file.
	aboutPageWASMPath = "/static/wasm/about.wasm"
)

var (
	// webFS is the embedded file system for the web directory.
	//go:embed web
	webFS embed.FS
)

// registerStaticFiles configures the server to serve static files
// from the embedded file system.
func (server *Server) registerStaticFiles() error {
	staticFiles, err := fs.Sub(webFS, staticFilesDirectory)
	if err != nil {
		return fmt.Errorf("failed to open static files directory error=%w", err)
	}
	server.router.Handle(staticFilesPathPrefix, http.StripPrefix(staticFilesPathPrefix, http.FileServerFS(staticFiles)))
	return nil
}

// registerPageRoutes configures the server to serve HTML pages.
func (server *Server) registerPageRoutes() error {
	return nil
}

// newPageData creates and returns a new pageData instance.
func newPageData(wasmPath string) *pageData {
	pageData := &pageData{
		Title:    pageTitleDefault,
		WASMPath: wasmPath,
	}
	return pageData
}
