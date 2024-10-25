// main.go
package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

var db *sql.DB

func main() {
	var err error
	// Connect to DuckDB (creates a new database file if not existing)
	db, err = sql.Open("duckdb", "auth.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Initialize database tables
	InitDatabase(db)

	// Retrieve and display all users (for testing purposes)
	users, err := GetAllUsers(db)
	if err != nil {
		log.Printf("Failed to retrieve users: %v", err)
	} else {
		fmt.Println("Users in the database:")
		for _, user := range users {
			fmt.Printf("UserID: %d, Username: %s\n", user.UserID, user.Username)
		}
	}

	// Set up HTTP routes
	http.HandleFunc("/register", RegisterHandler)
	http.HandleFunc("/login", LoginHandler)
	http.HandleFunc("/users", UsersHandler)

	// Start the server
	log.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
