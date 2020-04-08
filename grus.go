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

	version := "v0.3.0"

	// Print version if first argument is version.
	if os.Args[1] == "version" {
		fmt.Printf("Grus %s\n", version)
		os.Exit(0)
	}

	// Define default environment variables.
	envVar := make(map[string]bool)
	envVar["GRUS_SEARCH_ALL"] = false
	envVar["GRUS_ANAGRAMS"] = true
	envVar["GRUS_STRICT_UNJUMBLE"] = false
	envVar["GRUS_PRINT_PATH"] = false

	// Check environment variables.
	for k, _ := range envVar {
		env := os.Getenv(k)
		if env == "false" ||
			env == "0" {
			envVar[k] = false
		} else if env == "true" ||
			env == "1" {
			envVar[k] = true
		}
	}

	// Print environment variables if first argument is env.
	if os.Args[1] == "env" {
		for k, v := range envVar {
			fmt.Printf("%s: %t\n", k, v)
		}
		os.Exit(0)
	}

	// Define default dictionaries.
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

	// We use this to record if the word was unjumbled.
	unjumbled := false

	for k, dict := range dicts {
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
		if envVar["GRUS_PRINT_PATH"] {
			if k != 0 {
				fmt.Println()
			}
			fmt.Println(dict)
		}

		file, err := os.Open(dict)
		panicOnErr(err)

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
				if !envVar["GRUS_ANAGRAMS"] {
					os.Exit(0)
				}
				unjumbled = true
			}
		}
		panicOnErr(scanner.Err())
		file.Close()

		// If the user has asked to strictly unjumble then we
		// cannot exit till it's unjumbled.
		if envVar["GRUS_STRICT_UNJUMBLE"] &&
			!unjumbled {
			// If user has asked to strictly unjumble & we
			// haven't done that yet & this is the last
			// dictionary then we've failed to unjumble it
			// & the program must exit with a non-zero
			// exit code.
			if k == len(dicts)-1 {
				os.Exit(1)
			}

			// Cannot exit, must search next dictionary.
			continue
		}

		// If user hasn't asked to search all dictionaries
		// then exit the program.
		if !envVar["GRUS_SEARCH_ALL"] {
			os.Exit(0)
		}

	}
}

func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}
