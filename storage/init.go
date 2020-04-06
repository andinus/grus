package storage

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// initErr will log the error and close the database connection if
// necessary.
func initErr(db *DB, err error) {
	if db.Conn != nil {
		db.Conn.Close()
	}
	log.Fatalf("Initialization Error :: %s", err.Error())
}

func initDB(db *DB) {
	var err error

	db.Path = fmt.Sprintf("%s/grus.db", GetDir())

	db.Conn, err = sql.Open("sqlite3", db.Path)
	if err != nil {
		log.Printf("storage/init.go: %s\n",
			"Failed to open database connection")
		initErr(db, err)
	}

	sqlstmt := []string{
		`CREATE TABLE IF NOT EXISTS words (
        word   TEXT PRIMARY KEY NOT NULL,
        sorted TEXT NOT NULL);`,
		`INSERT INTO words(word, lexical)
        values("grus", "grsu");`,
	}

	// We range over statements and execute them one by one, this
	// is during initialization so it doesn't matter if it takes
	// few more ms. This way we know which statement caused the
	// program to fail.
	for _, s := range sqlstmt {
		stmt, err := db.Conn.Prepare(s)

		if err != nil {
			log.Printf("storage/init.go: %s\n",
				"failed to prepare statement")
			log.Println(s)
			initErr(db, err)
		}

		_, err = stmt.Exec()
		stmt.Close()
		if err != nil {
			log.Printf("storage/init.go: %s\n",
				"failed to execute statement")
			log.Println(s)
			initErr(db, err)
		}
	}
}
