package server

type (
	pageData struct {
		Title    string
		WASMPath string
	}
)

const (
	pageDataTitleDefault = "DummyAI"
	homePagePath         = "GET /"
	homePageWASMPath     = "/wasm/home.wasm"
	aboutPagePath        = "GET /about"
	aboutPageWASMPath    = "/wasm/about.wasm"
)

var (
	homePageData  = newPageData(homePageWASMPath)
	aboutPageData = newPageData(aboutPageWASMPath)

	pageRoutes = map[string]*pageData{
		homePagePath:  homePageData,
		aboutPagePath: aboutPageData,
	}
)

func newPageData(wasmPath string) *pageData {
	pageData := &pageData{
		Title:    pageDataTitleDefault,
		WASMPath: wasmPath,
	}
	return pageData
}
