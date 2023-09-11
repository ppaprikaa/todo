package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/sanzhh/todo/internal/lib/e"
)

func NewSQLite(filepath string) (_ *sql.DB, err error) {
	var (
		errorMessage = "sqlite connection failed"
	)

	defer func() { err = e.WrapS(errorMessage, err) }()

	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
