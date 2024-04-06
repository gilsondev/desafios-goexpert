package database

import (
	"context"
	"log"
	"time"
)

func CreateSchema() error {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Millisecond)
	defer cancel()

	db, err := OpenConnection()
	if err != nil {
		log.Fatalf("Failed to open connection with database: %v", err)
		panic(err)
	}
	defer db.Close()

	_, err = db.ExecContext(ctx, `
		CREATE TABLE IF NOT EXISTS exchange (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			code TEXT NOT NULL,
			codein TEXT NOT NULL,
			name TEXT NOT NULL,
			high REAL NOT NULL,
			low REAL NOT NULL,
			var_bid REAL NOT NULL,
			pct_change REAL NOT NULL,
			bid REAL NOT NULL,
			ask REAL NOT NULL,
			timestamp INTEGER NOT NULL,
			create_date TEXT NOT NULL
		);
	`)
	if err != nil {
		log.Fatalf("Failed to create schema: %v", err)
		return err
	}

	return nil
}
