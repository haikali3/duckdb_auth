package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/marcboeker/go-duckdb"
)

func main() {
	// Connect to DuckDB (creates a new database file if not existing)
	db, err := sql.Open("duckdb", "auth.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create tables for users and roles
	createTables := `
    CREATE TABLE IF NOT EXISTS users (
        user_id INTEGER PRIMARY KEY AUTOINCREMENT,
        username VARCHAR UNIQUE,
        password_hash VARCHAR
    );
    CREATE TABLE IF NOT EXISTS roles (
        role_id INTEGER PRIMARY KEY AUTOINCREMENT,
        role_name VARCHAR UNIQUE
    );
    CREATE TABLE IF NOT EXISTS user_roles (
        user_id INTEGER,
        role_id INTEGER,
        FOREIGN KEY (user_id) REFERENCES users(user_id),
        FOREIGN KEY (role_id) REFERENCES roles(role_id)
    );
    `

	if _, err := db.Exec(createTables); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database initialized with users, roles, and user_roles tables.")
}
