package lexical

import "testing"

// TestSort tests the Sort func.
func TestSort(t *testing.T) {
	words := make(map[string]string)

	words["dcba"] = "abcd"
	words["zyx"] = "xyz"

	for word, sorted := range words {
		s := Sort(word)
		if s != sorted {
			t.Errorf("Sort func failed, got %s, want %s",
				s, sorted)
		}
	}
}
