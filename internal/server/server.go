package server

type (
	pageData struct {
		Title    string
		WASMPath string
	}
)

const (
	titleDefault     = "DummyAI"
	homePath         = "GET /"
	homeWASMPath     = "/wasm/home.wasm"
	aboutPath        = "GET /about"
	aboutWASMPath    = "/wasm/about.wasm"
	catchAllPath     = "/"
	error404WASMPath = "/wasm/error_404.wasm"
	error500WASMPath = "/wasm/error_500.wasm"
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

	error404PageData = &pageData{
		Title:    titleDefault,
		WASMPath: error404WASMPath,
	}

	error500PageData = &pageData{
		Title:    titleDefault,
		WASMPath: error500WASMPath,
	}
)
