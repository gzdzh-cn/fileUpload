package config

import (
	"dzhgo/addons/fileUpload"
	"github.com/gzdzh/dzhcore"
)

func init() {
	dzhcore.SetVersions("fileUpload", fileUpload.Version)
}
