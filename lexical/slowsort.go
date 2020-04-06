package lexical

import (
	"sort"
	"strings"
)

// SlowSort returns string in lexical order. This function is slower
// than Lexical.
func SlowSort(word string) (sorted string) {
	// Convert word to a slice, sort the slice.
	t := strings.Split(word, "")
	sort.Strings(t)

	sorted = strings.Join(t, "")
	return
}
