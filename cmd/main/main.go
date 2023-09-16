package main

import (
	"log"

	"github.com/ppaprikaa/todo/internal/commands"
	"github.com/ppaprikaa/todo/internal/db"
)

func main() {
	db, err := db.NewSQLite()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	commands.Execute()
}
