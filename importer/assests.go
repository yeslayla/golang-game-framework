package importer

import (
	"embed"
)

var EmbededAssets *embed.FS = nil

func SetAssets(assets *embed.FS) {
	EmbededAssets = assets
}
