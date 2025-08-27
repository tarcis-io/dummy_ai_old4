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

const (
	pageTitleDefault = "DummyAI"
)

func New(address string) (*Server, error) {
	server := &Server{
		address: address,
	}
	return server, nil
}

func newPageData(wasmPath string) *pageData {
	pageData := &pageData{
		Title:    pageTitleDefault,
		WASMPath: wasmPath,
	}
	return pageData
}
