package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"tildegit.org/andinus/grus/lexical"
	"tildegit.org/andinus/grus/search"
	"tildegit.org/andinus/grus/storage"
)

func grus() {
	version := "v0.1.0"

	// Early Check: If command was not passed then print usage and
	// exit. Later command & service both are checked, this check
	// is for version command. If not checked then running grus
	// without any args will fail because os.Args[1] will panic
	// the program & produce runtime error.
	if len(os.Args) == 1 || len(os.Args[1]) == 0 {
		printUsage()
		os.Exit(0)
	}

	// Running just `grus` would've paniced the program here if
	// length of os.Args was not checked beforehand because there
	// would be no os.Args[1].
	switch os.Args[1] {
	case "version", "v", "-version", "--version", "-v":
		fmt.Printf("Grus %s\n", version)
		os.Exit(0)
	case "help", "-help", "--help", "-h":
		printUsage()
		os.Exit(0)
	case "init", "i":
		db := storage.Init()
		db.Conn.Close()
		os.Exit(0)
	}

	// Initialize the database connection.
	db := storage.InitConn()
	defer db.Conn.Close()

	word := os.Args[1]
	sorted := lexical.Sort(word)

	anagrams, err := search.Anagrams(sorted, db)
	if err == sql.ErrNoRows {
		fmt.Println("Word not found in database.")
		return
	} else if err != nil {
		log.Fatalf("grus: Search failed :: %s", err)
	}
	for _, w := range anagrams {
		fmt.Println(w)
	}
}
