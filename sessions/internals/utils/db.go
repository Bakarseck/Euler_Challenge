package utils

import (
	"database/sql"
	"log"
)

func OpenDatabase(dbPath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func ExecuteQuery(db *sql.DB, query string) error {
	_, err := db.Exec(query)
	if err != nil {
		log.Printf("Error executing query: %s", err)
		return err
	}
	return nil
}
