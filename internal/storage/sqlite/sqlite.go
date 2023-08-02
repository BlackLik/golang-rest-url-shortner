package sqlite

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	db *sql.DB
}

func New(storage_path string) (*Storage, error) {

	const op = "storage.sqlite.New"
	db, err := sql.Open("sqlite3", storage_path)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	stmt, err := db.Prepare(`
		CREATE TABLE IF NOT EXISTS urls (
			id INTEGER PRIMARY KEY AUTOINCREMENT, 
			alias TEXT NOT NULL UNIQUE, 
			url TEXT NOT NULL, 
            created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
		);
		CREATE INDEX IF NOT EXISTS idx_urls_alias ON urls(alias);	
	`)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.Exec()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{db: db}, nil
}
