package server

import (
	"embed"
	"net/http"
	"text/template"
)

type (
	Server struct {
		address      string
		serveMux     *http.ServeMux
		pageTemplate *template.Template
	}

	pageData struct {
		Title    string
		WASMPath string
	}
)

func (server *Server) RegisterHandler(route string, handler func(http.ResponseWriter, *http.Request)) {
	server.serveMux.HandleFunc(route, handler)
}

func (server *Server) ListenAndServe() error {
	return http.ListenAndServe(server.address, server.serveMux)
}

func (server *Server) homeHandler(responseWriter http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/" {
		server.error404Handler(responseWriter, request)
		return
	}
	err := server.renderPage(responseWriter, homePageData)
	if err != nil {
		server.error500Handler(responseWriter, request)
	}
}

func (server *Server) aboutHandler(responseWriter http.ResponseWriter, request *http.Request) {
	err := server.renderPage(responseWriter, aboutPageData)
	if err != nil {
		server.error500Handler(responseWriter, request)
	}
}

func (server *Server) error404Handler(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.WriteHeader(http.StatusNotFound)
	err := server.renderPage(responseWriter, error404PageData)
	if err != nil {
		server.error500Handler(responseWriter, request)
	}
}

func (server *Server) error500Handler(responseWriter http.ResponseWriter, _ *http.Request) {
	responseWriter.WriteHeader(http.StatusInternalServerError)
	err := server.renderPage(responseWriter, error500PageData)
	if err != nil {
		http.Error(responseWriter, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func (server *Server) renderPage(responseWriter http.ResponseWriter, pageData *pageData) error {
	return server.pageTemplate.Execute(responseWriter, pageData)
}

const (
	titleDefault = "DummyAI"
)

var (
	homePageData = &pageData{
		Title:    titleDefault,
		WASMPath: "/wasm/home.wasm",
	}

	aboutPageData = &pageData{
		Title:    titleDefault,
		WASMPath: "/wasm/about.wasm",
	}

	error404PageData = &pageData{
		Title:    titleDefault,
		WASMPath: "/wasm/error_404.wasm",
	}

	error500PageData = &pageData{
		Title:    titleDefault,
		WASMPath: "/wasm/error_500.wasm",
	}

	//go:embed web/template/*.html
	pageTemplateFS embed.FS

	pageTemplate = template.Must(template.ParseFS(pageTemplateFS, "web/template/*.html"))
)

func New(address string) *Server {
	server := &Server{
		address:      address,
		serveMux:     http.NewServeMux(),
		pageTemplate: pageTemplate,
	}
	server.RegisterHandler("/", server.homeHandler)
	server.RegisterHandler("/about", server.aboutHandler)
	return server
}
