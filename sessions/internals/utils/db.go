package utils

import (
	"database/sql"
	"log"
	"os"
)

func OpenDatabase(dbPath string, sqlFile string) (*sql.DB, error) {
	// Ouvrir ou créer la base de données
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	if _, err := os.Stat(sqlFile); err == nil {
		sqlContent, err := os.ReadFile(sqlFile)
		if err != nil {
			return nil, err
		}

		if _, err := db.Exec(string(sqlContent)); err != nil {
			return nil, err
		}
	} else {
		log.Printf("Le fichier SQL '%s' n'a pas été trouvé. Aucune table n'a été créée.", sqlFile)
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
