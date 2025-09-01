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
	return server, nil
}

func (server *Server) Start() error {
	err := http.ListenAndServe(server.address, server.router)
	if err != nil {
		return fmt.Errorf("failed to start server address=%s error=%w", server.address, err)
	}
	return nil
}
