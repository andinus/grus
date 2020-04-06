package search

import "tildegit.org/andinus/grus/storage"

// Anagrams will search for unjumbled words in database, given sorted
// word along with all the anagrams.
func Anagrams(sorted string, db *storage.DB) (anagrams []string, err error) {
	db.Mu.RLock()
	defer db.Mu.RUnlock()

	stmt, err := db.Conn.Prepare("SELECT word FROM words WHERE sorted = ?")
	if err != nil {
		return
	}
	defer stmt.Close()

	rows, err := stmt.Query(sorted)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var word string
		err = rows.Scan(&word)
		if err != nil {
			return
		}
		anagrams = append(anagrams, word)
	}
	err = rows.Err()
	if err != nil {
		return
	}
	return
}
