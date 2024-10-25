// main.go
package main

import (
	"database/sql"
	"fmt"
	"log"
)

func main() {
	// Connect to DuckDB (creates a new database file if not existing)
	db, err := sql.Open("duckdb", "auth.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Initialize database tables
	InitDatabase(db)

	// Test registering a user
	username := "testuser"
	password := "securepassword"

	err = RegisterUser(db, username, password)
	if err != nil {
		log.Printf("User registration failed: %v", err)
	} else {
		fmt.Println("User registered successfully!")
	}

	// Test user authentication
	success, err := AuthenticateUser(db, username, password)
	if err != nil {
		log.Fatal(err)
	}
	if success {
		fmt.Println("User authenticated successfully!")
	} else {
		fmt.Println("Authentication failed.")
	}
}
