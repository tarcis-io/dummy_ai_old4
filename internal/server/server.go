package server

type (
	Server struct {
		address string
	}

	pageData struct {
		Title    string
		WASMPath string
	}
)

func New(address string) (*Server, error) {
	server := &Server{
		address: address,
	}
	return server, nil
}

func newPageData(wasmPath string) *pageData {
	pageData := &pageData{
		WASMPath: wasmPath,
	}
	return pageData
}
