// +build linux netbsd openbsd freebsd dragonfly

package storage

import (
	"fmt"
	"os"
)

// GetDir returns grus data directory. Check if the user has set
// GRUS_DIR, if not then check if XDG_DATA_HOME is set & if that is
// not set then assume it to be the default value which is
// $HOME/.local/share according to XDG Base Directory Specification.
func GetDir() string {
	cacheDir := SysDir()

	// Grus cache directory is cacheDir/grus.
	grusCacheDir := fmt.Sprintf("%s/%s", cacheDir,
		"grus")

	return grusCacheDir
}

// SysDir returns the system data directory, this is useful for unveil in
// OpenBSD.
func SysDir() string {
	cacheDir := os.Getenv("GRUS_DIR")
	if len(cacheDir) == 0 {
		cacheDir = os.Getenv("XDG_DATA_HOME")
	}
	if len(cacheDir) == 0 {
		cacheDir = fmt.Sprintf("%s/%s/%s", os.Getenv("HOME"),
			".local", "share")
	}

	return cacheDir
}
