package mock

import (
	"os"

	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/rai-project/dlframework"
	"github.com/rai-project/dlframework/framework"
)

// FrameworkManifest ...
var FrameworkManifest = dlframework.FrameworkManifest{
	Name:    "mock",
	Version: "1.0.0",
}

func assetFS() *assetfs.AssetFS {
	assetInfo := func(path string) (os.FileInfo, error) {
		return os.Stat(path)
	}
	for k := range _bintree.Children {
		return &assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: assetInfo, Prefix: k}
	}
	panic("unreachable")
}

func Register() {
	if supportedSystem {
		framework.Register(FrameworkManifest, assetFS())
	}
}
