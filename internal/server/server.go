package server

type (
	pageData struct {
		Title    string
		WASMPath string
	}
)

const (
	pageDataTitleDefault = "DummyAI"
)

func newPageData(wasmPath string) *pageData {
	pageData := &pageData{
		Title:    pageDataTitleDefault,
		WASMPath: wasmPath,
	}
	return pageData
}
