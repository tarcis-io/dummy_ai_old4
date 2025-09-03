// Package server provides a pre-configured HTTP server for the application.
package server

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"
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

	// pageHeaderContentTypeKey is the header key for the content type.
	pageHeaderContentTypeKey = "Content-Type"

	// pageHeaderContentTypeValue is the header value for the content type.
	pageHeaderContentTypeValue = "text/html; charset=UTF-8"

	pageHeaderContentSecurityPolicyKey = "Content-Security-Policy"

	pageHeaderContentSecurityPolicyValue = "default-src 'self';"

	pageHeaderXContentTypeOptionsKey = "X-Content-Type-Options"

	pageHeaderXContentTypeOptionsValue = "nosniff"

	pageHeaderXFrameOptionsKey = "X-Frame-Options"

	pageHeaderXFrameOptionsValue = "DENY"

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

// New creates, configures and returns a new Server instance.
func New(address string) (*Server, error) {
	server := &Server{
		address: address,
		router:  http.NewServeMux(),
	}
	err := server.registerStaticFiles()
	if err != nil {
		return nil, err
	}
	err = server.registerPageRoutes()
	if err != nil {
		return nil, err
	}
	return server, nil
}

// Start starts the server and listens for incoming requests.
func (server *Server) Start() error {
	err := http.ListenAndServe(server.address, server.router)
	if err != nil {
		return fmt.Errorf("failed to start server address=%s error=%w", server.address, err)
	}
	return nil
}

// registerStaticFiles configures the server to serve static files
// from the embedded file system.
// It returns an error if the static files directory cannot be opened.
func (server *Server) registerStaticFiles() error {
	staticFiles, err := fs.Sub(webFS, staticFilesDirectory)
	if err != nil {
		return fmt.Errorf("failed to open static files directory error=%w", err)
	}
	server.router.Handle(staticFilesPathPrefix, http.StripPrefix(staticFilesPathPrefix, http.FileServerFS(staticFiles)))
	return nil
}

// registerPageRoutes configures the server to serve HTML pages.
// It returns an error if the page template cannot be parsed or if a page cannot be registered.
func (server *Server) registerPageRoutes() error {
	pageTemplate, err := template.ParseFS(webFS, pageTemplatePattern)
	if err != nil {
		return fmt.Errorf("failed to parse page template error=%w", err)
	}
	pageHeaders := map[string]string{
		pageHeaderContentTypeKey:           pageHeaderContentTypeValue,
		pageHeaderContentSecurityPolicyKey: pageHeaderContentSecurityPolicyValue,
		pageHeaderXContentTypeOptionsKey:   pageHeaderXContentTypeOptionsValue,
		pageHeaderXFrameOptionsKey:         pageHeaderXFrameOptionsValue,
	}
	pageRoutes := map[string]*pageData{
		homePagePath:  newPageData(homePageWASMPath),
		aboutPagePath: newPageData(aboutPageWASMPath),
	}
	for pagePath, pageData := range pageRoutes {
		pageBuffer := new(bytes.Buffer)
		err = pageTemplate.Execute(pageBuffer, pageData)
		if err != nil {
			return fmt.Errorf("failed to execute page template error=%w", err)
		}
		pageCache := pageBuffer.Bytes()
		server.router.HandleFunc(pagePath, func(responseWriter http.ResponseWriter, request *http.Request) {
			for pageHeaderKey, pageHeaderValue := range pageHeaders {
				responseWriter.Header().Set(pageHeaderKey, pageHeaderValue)
			}
			responseWriter.Write(pageCache)
		})
	}
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
