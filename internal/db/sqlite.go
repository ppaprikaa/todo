package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/sanzhh/todo/internal/config"
	"github.com/sanzhh/todo/internal/lib/e"
)

func NewSQLite() (_ *sql.DB, err error) {
	var (
		errorMessage = "sqlite connection failed"
	)

	defer func() { err = e.WrapS(errorMessage, err) }()

	db, err := sql.Open("sqlite3", config.CFG.DSN)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	query := `CREATE TABLE IF NOT EXISTS todos (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		date TEXT NOT NULL,
		done BOOLEAN NOT NULL DEFAULT false
	)`

	if _, err = db.Exec(query); err != nil {
		return nil, err
	}

	return db, nil
}
