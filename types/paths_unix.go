package types

import (
	"os"
	"path/filepath"
)

// ConfigPath holds the path for hoist's config
// files.
var ConfigPath string

func init() {
	homepath := os.Getenv("HOME")
	ConfigPath = filepath.Join(homepath, ".config", "hoist")
}
