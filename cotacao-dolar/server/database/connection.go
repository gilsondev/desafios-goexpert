package database

import (
	"database/sql"

	_ "github.com/glebarez/go-sqlite"
)

func OpenConnection() (*sql.DB, error) {
	db, err := sql.Open("sqlite", "cotacao-dolar.sqlite")
	if err != nil {
		return nil, err
	}

	return db, nil
}
