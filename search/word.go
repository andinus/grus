package search

import "tildegit.org/andinus/grus/storage"

// Word will search for unjumbled words in database, given sorted word.
func Word(sorted string, db *storage.DB) (word string, err error) {
	db.Mu.RLock()
	defer db.Mu.RUnlock()

	stmt, err := db.Conn.Prepare("SELECT word FROM words WHERE sorted = ?")
	if err != nil {
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(sorted).Scan(&word)
	if err != nil {
		return
	}
	return
}
