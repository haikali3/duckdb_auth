// db.go
package main

import (
	"database/sql"
	_ "github.com/marcboeker/go-duckdb"
	"log"
)

// InitDatabase creates the necessary tables for authentication
func InitDatabase(db *sql.DB) {
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
	log.Println("Database initialized with users, roles, and user_roles tables.")
}
