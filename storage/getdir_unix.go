// +build linux netbsd openbsd freebsd dragonfly

package storage

import (
	"fmt"
	"os"
)

// GetDir returns grus data directory. Check if the user has set
// XDG_DATA_HOME is set & if that is not set then assume it to be the
// default value which is $HOME/.local/share according to XDG Base
// Directory Specification.
func GetDir() (grusCacheDir string) {
	cacheDir := SysDir()

	// Grus cache directory is cacheDir/grus.
	grusCacheDir = fmt.Sprintf("%s/%s", cacheDir,
		"grus")

	return
}

// SysDir returns the system data directory, this is useful for unveil in
// OpenBSD.
func SysDir() (cacheDir string) {
	cacheDir = os.Getenv("XDG_DATA_HOME")
	if len(cacheDir) == 0 {
		cacheDir = fmt.Sprintf("%s/%s/%s", os.Getenv("HOME"),
			".local", "share")
	}

	return
}
