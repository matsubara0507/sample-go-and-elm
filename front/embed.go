package front

import _ "embed"

//go:embed index.html
var indexHtmlFile []byte

func EmbedIndexHtml() []byte {
	return indexHtmlFile
}
