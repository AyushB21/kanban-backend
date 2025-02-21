package database

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite" // ✅ Use modernc.org/sqlite instead of go-sqlite3
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite", "kanban.db") // ✅ Use "sqlite" instead of "sqlite3"
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Create tasks table
	createTasksTableQuery := `
	CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		description TEXT,
		status TEXT DEFAULT 'pending',
		column TEXT DEFAULT 'To Do',
		assignee TEXT
	);
	`
	_, err = DB.Exec(createTasksTableQuery)
	if err != nil {
		log.Fatal("Failed to create tasks table:", err)
	}

	// ✅ Add users table for authentication
	createUsersTableQuery := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT UNIQUE NOT NULL,
		password TEXT NOT NULL
	);
	`
	_, err = DB.Exec(createUsersTableQuery)
	if err != nil {
		log.Fatal("Failed to create users table:", err)
	}

	log.Println("Connected to SQLite database and ensured required tables exist!")
}
