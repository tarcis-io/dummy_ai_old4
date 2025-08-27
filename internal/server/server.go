package server

type (
	pageData struct {
		Title    string
		WASMPath string
	}
)

const (
	pageTitleDefault = "DummyAI"
)

func newPageData(wasmPath string) *pageData {
	pageData := &pageData{
		Title:    pageTitleDefault,
		WASMPath: wasmPath,
	}
	return pageData
}
