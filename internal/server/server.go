package server

type (
	pageData struct {
		Title    string
		WASMPath string
	}
)

const (
	pageTitleDefault  = "DummyAI"
	homePagePath      = "GET /"
	homePageWASMPath  = "/static/wasm/home.wasm"
	aboutPagePath     = "GET /about"
	aboutPageWASMPath = "/static/wasm/about.wasm"
)

var (
	homePageData  = newPageData(homePageWASMPath)
	aboutPageData = newPageData(aboutPageWASMPath)
)

func newPageData(wasmPath string) *pageData {
	pageData := &pageData{
		Title:    pageTitleDefault,
		WASMPath: wasmPath,
	}
	return pageData
}
