package admin

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/qor/assetfs"
)

var (
	root, _         = os.Getwd()
	globalViewPaths []string
	globalAssetFSes []assetfs.Interface
)

func init() {
	if path := os.Getenv("WEB_ROOT"); path != "" {
		root = path
	}
}

// RegisterViewPath register view path for all assetfs
func RegisterViewPath(pth string) {
	globalViewPaths = append(globalViewPaths, pth)

	for _, assetFS := range globalAssetFSes {
		if assetFS.RegisterPath(filepath.Join(root, "vendor", pth)) != nil {
			for _, gopath := range strings.Split(os.Getenv("GOPATH"), ":") {
				if assetFS.RegisterPath(filepath.Join(gopath, "src", pth)) == nil {
					break
				}
			}
		}
	}
}
