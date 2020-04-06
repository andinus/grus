package storage

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
)

// DB holds the database connection, mutex & path.
type DB struct {
	Path string
	Mu   *sync.RWMutex
	Conn *sql.DB
}

// Init initializes the database.
func Init() *DB {
	db := DB{
		Mu: new(sync.RWMutex),
	}

	initDB(&db)
	return &db
}

// InitConn initializes database connection.
func InitConn() *DB {
	var err error
	db := DB{
		Mu: new(sync.RWMutex),
	}

	db.Path = fmt.Sprintf("%s/grus.db", GetDir())

	db.Conn, err = sql.Open("sqlite3", db.Path)
	if err != nil {
		log.Printf("storage/init.go: %s\n",
			"Failed to open database connection")
		initErr(&db, err)
	}

	return &db
}
