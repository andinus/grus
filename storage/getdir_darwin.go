// +build darwin

package storage

import (
	"fmt"
	"os"
)

// GetDir returns grus data directory. Default data directory on
// macOS is $HOME/Library.
func GetDir() string {
	cacheDir := fmt.Sprintf("%s/%s",
		os.Getenv("HOME"),
		"Library")

	// Grus cache directory is cacheDir/grus
	grusCacheDir := fmt.Sprintf("%s/%s", cacheDir,
		"grus")

	return grusCacheDir
}
