package main

import (
	"log"

	"github.com/sanzhh/todo/internal/commands"
	"github.com/sanzhh/todo/internal/db"
)

func main() {
	db, err := db.NewSQLite()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	commands.Execute()
}
