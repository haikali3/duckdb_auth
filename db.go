// db.go
package main

import (
	"database/sql"
	"log"

	_ "github.com/marcboeker/go-duckdb"
)

// InitDatabase creates the necessary tables and sequence for authentication
func InitDatabase(db *sql.DB) {
	createTables := `
    CREATE SEQUENCE IF NOT EXISTS user_id_seq;
    CREATE TABLE IF NOT EXISTS users (
        user_id INTEGER PRIMARY KEY,
        username VARCHAR UNIQUE,
        password_hash VARCHAR
    );
    CREATE TABLE IF NOT EXISTS roles (
        role_id INTEGER PRIMARY KEY,
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
