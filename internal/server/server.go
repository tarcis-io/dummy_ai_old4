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
	pageTitleDefault  = "DummyAI"
	homePageWASMPath  = "/static/wasm/home.wasm"
	aboutPageWASMPath = "/static/wasm/about.wasm"
)

var (
	homePageData  = newPageData(homePageWASMPath)
	aboutPageData = newPageData(aboutPageWASMPath)
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
