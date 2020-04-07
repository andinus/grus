// +build openbsd

package main

import (
	"os"

	"golang.org/x/sys/unix"
	"tildegit.org/andinus/lynx"
)

func main() {
	err := unix.PledgePromises("unveil stdio rpath")
	panicOnErr(err)

	unveil()

	// Drop unveil from promises.
	err = unix.PledgePromises("stdio rpath")
	panicOnErr(err)

	grus()
}

func unveil() {
	paths := make(map[string]string)

	paths["/usr/share/dict"] = "r"
	paths["/usr/local/share/dict"] = "r"

	// Unveil user defined dictionaries.
	if len(os.Args) >= 3 {
		for _, dict := range os.Args[2:] {
			paths[dict] = "r"
		}
	}
	// This will not return error if the file doesn't exist.
	err := lynx.UnveilPaths(paths)
	panicOnErr(err)

	// Block further unveil calls.
	err = lynx.UnveilBlock()
	panicOnErr(err)
}
