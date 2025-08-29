package server

type (
	Server struct {
		address string
	}
)

func New(address string) (*Server, error) {
	server := &Server{
		address: address,
	}
	return server, nil
}
