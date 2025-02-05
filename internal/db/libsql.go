package db

import (
	"database/sql"
	"log"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("libsql", "file:heal.db")
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Example table creation
	createTable := `
	CREATE TABLE IF NOT EXISTS content (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		body TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`
	_, err = DB.Exec(createTable)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}
}
