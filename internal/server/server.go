package server

type (
	pageData struct {
		Title    string
		WASMPath string
	}
)

const (
	pageDataTitleDefault  = "DummyAI"
	homePageRoutePath     = "GET /"
	homePageDataWASMPath  = "/wasm/home.wasm"
	aboutPageRoutePath    = "GET /about"
	aboutPageDataWASMPath = "/wasm/about.wasm"
)

var (
	homePageData  = newPageData(homePageDataWASMPath)
	aboutPageData = newPageData(aboutPageDataWASMPath)

	pageRoutes = map[string]*pageData{
		homePageRoutePath:  homePageData,
		aboutPageRoutePath: aboutPageData,
	}
)

func newPageData(wasmPath string) *pageData {
	pageData := &pageData{
		Title:    pageDataTitleDefault,
		WASMPath: wasmPath,
	}
	return pageData
}
