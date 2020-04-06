package lexical

import "sort"

// Sort takes a string as input and returns the lexical order.
func Sort(word string) (sorted string) {
	// Convert the string to []rune.
	var r []rune
	for _, char := range word {
		r = append(r, char)
	}

	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})

	sorted = string(r)
	return
}
