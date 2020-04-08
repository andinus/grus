package main

import (
	"bufio"
	"fmt"
	"os"

	"tildegit.org/andinus/grus/lexical"
)

func grus() {
	if len(os.Args) == 1 {
		fmt.Println("Usage: grus <word> <dictionaries>")
		os.Exit(1)
	}

	version := "v0.2.1"

	if os.Args[1] == "version" {
		fmt.Printf("Grus %s\n", version)
		os.Exit(0)
	}

	dicts := []string{
		"/usr/local/share/dict/words",
		"/usr/local/share/dict/web2",
		"/usr/share/dict/words",
		"/usr/share/dict/web2",
		"/usr/share/dict/special/4bsd",
		"/usr/share/dict/special/math",
	}

	// User has specified dictionaries, prepend them to dicts
	// list.
	if len(os.Args) >= 3 {
		dicts = append(os.Args[2:], dicts...)
	}

	// Check if user has asked to search in all dictionaries.
	searchAll := false
	searchAllEnv := os.Getenv("GRUS_SEARCH_ALL")
	if searchAllEnv == "true" ||
		searchAllEnv == "1" {
		searchAll = true
	}

	// Check if user wants anagrams.
	anagrams := false
	anagramsEnv := os.Getenv("GRUS_ANAGRAMS")
	if anagramsEnv == "true" ||
		anagramsEnv == "1" {
		anagrams = true
	}

	// Check if user wants to print dictionary path.
	printPath := false
	printPathEnv := os.Getenv("GRUS_PRINT_PATH")
	if printPathEnv == "true" ||
		printPathEnv == "1" {
		printPath = true
	}

	for _, dict := range dicts {
		if _, err := os.Stat(dict); err != nil &&
			!os.IsNotExist(err) {
			// Error is not nil & also it's not path
			// doesn't exist error. We do it this way to
			// avoid another level of indentation.
			panic(err)
		} else if err != nil &&
			os.IsNotExist(err) {
			// If file doesn't exist then continue with
			// next dictionary.
			continue
		}

		// Print path to dictionary if printPath is true.
		if printPath {
			fmt.Println(dict)
		}

		file, err := os.Open(dict)
		panicOnErr(err)
		defer file.Close()

		// We use this to record if the word was unjumbled.
		unjumbled := false

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			// Filter words by comparing length first &
			// run lexical.Sort on it only if the length
			// is equal.
			if len(scanner.Text()) == len(os.Args[1]) &&
				lexical.Sort(scanner.Text()) == lexical.Sort(os.Args[1]) {
				fmt.Println(scanner.Text())
				// If the user doesn't want anagrams
				// then exit the program.
				if !anagrams {
					os.Exit(0)
				}
				unjumbled = true
			}
		}
		panicOnErr(scanner.Err())

		// If word was unjumbled & user hasn't asked to search
		// in all dictionaries then exit the program otherwise
		// keep searching in other dictionaries.
		if unjumbled &&
			!searchAll {
			os.Exit(0)
		}
	}
}

func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}
