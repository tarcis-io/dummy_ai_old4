package server

import (
	"fmt"
	"net/http"
)

type (
	Server struct {
		address string
		router  *http.ServeMux
	}
)

func New(address string) (*Server, error) {
	server := &Server{
		address: address,
		router:  http.NewServeMux(),
	}
	err := server.registerStaticFiles()
	if err != nil {
		return nil, fmt.Errorf("failed to register static files error=%w", err)
	}
	err = server.registerRoutes()
	if err != nil {
		return nil, fmt.Errorf("failed to register routes error=%w", err)
	}
	return server, nil
}

func (server *Server) Start() error {
	return http.ListenAndServe(server.address, server.router)
}

func (server *Server) registerStaticFiles() error {
	return nil
}

func (server *Server) registerRoutes() error {
	return nil
}
