// +build openbsd

package main

import (
	"log"
	"os"

	"golang.org/x/sys/unix"
	"tildegit.org/andinus/grus/storage"
	"tildegit.org/andinus/lynx"
)

func main() {
	unveil()
	grus()
}

func unveil() {
	path := storage.GetDir()
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		log.Fatalf("Unable to create directory: %s", path)
	}

	paths := make(map[string]string)

	paths[path] = "rwc"

	err = lynx.UnveilPathsStrict(paths)
	if err != nil {
		log.Fatal(err)
	}

	// Block further unveil calls.
	err = unix.UnveilBlock()
	if err != nil {
		log.Fatal(err)
	}
}
