package server

import (
	"net/http"
)

type (
	Server struct {
		address string
		handler *http.ServeMux
	}
)

func (s *Server) ListenAndServe() error {
	return http.ListenAndServe(s.address, s.handler)
}
