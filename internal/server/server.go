package server

type (
	pageData struct {
		Title    string
		WASMPath string
	}
)

const (
	titleDefault  = "DummyAI"
	homeWASMPath  = "/wasm/home.wasm"
	aboutWASMPath = "/wasm/about.wasm"
)

var (
	homePageData = &pageData{
		Title:    titleDefault,
		WASMPath: homeWASMPath,
	}

	aboutPageData = &pageData{
		Title:    titleDefault,
		WASMPath: aboutWASMPath,
	}
)
