package server

type (
	pageData struct {
		Title    string
		WASMPath string
	}
)

const (
	pageDataTitleDefault  = "DummyAI"
	homePagePath          = "GET /"
	homePageDataWASMPath  = "/wasm/home.wasm"
	aboutPagePath         = "GET /about"
	aboutPageDataWASMPath = "/wasm/about.wasm"
)

var (
	homePageData  = newPageData(homePageDataWASMPath)
	aboutPageData = newPageData(aboutPageDataWASMPath)

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
