package types

import (
	"os"
	"path/filepath"
)

var (
	// ConfigDir holds the path for hoist's config
	// files.
	ConfigDir string

	// ConfigFilePath holds the path for hoist's main
	// configuration file.
	ConfigFilePath string
)

func init() {
	homepath := os.Getenv("HOME")
	ConfigDir = filepath.Join(homepath, ".config", "hoist")
	ConfigFilePath = filepath.Join(ConfigDir, "config.json")
}
