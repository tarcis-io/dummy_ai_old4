package server

type (
	Server struct {
	}

	pageData struct {
		Title    string
		WASMPath string
	}
)

const (
	titleDefault = "DummyAI"
)

var (
	homePageData = &pageData{
		Title:    titleDefault,
		WASMPath: "/wasm/home.wasm",
	}

	aboutPageData = &pageData{
		Title:    titleDefault,
		WASMPath: "/wasm/about.wasm",
	}

	error404PageData = &pageData{
		Title:    titleDefault,
		WASMPath: "/wasm/error_404.wasm",
	}

	error500PageData = &pageData{
		Title:    titleDefault,
		WASMPath: "/wasm/error_500.wasm",
	}
)
