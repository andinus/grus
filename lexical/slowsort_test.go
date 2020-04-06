package lexical

import "testing"

// TestSlowSort tests the SlowSort func.
func TestSlowSort(t *testing.T) {
	words := make(map[string]string)

	words["dcba"] = "abcd"
	words["zyx"] = "xyz"

	for word, sorted := range words {
		s := SlowSort(word)
		if s != sorted {
			t.Errorf("Sort func failed, got %s, want %s",
				s, sorted)
		}
	}
}
