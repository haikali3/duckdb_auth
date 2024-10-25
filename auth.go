// auth.go
package main

import (
	"database/sql"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes a plain password using bcrypt
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPasswordHash compares a hashed password with a plain password
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// RegisterUser registers a new user in the DuckDB database
func RegisterUser(db *sql.DB, username, password string) error {
	hashedPassword, err := HashPassword(password)
	if err != nil {
		return err
	}

	// Fetch the next user_id from the sequence
	var userID int
	err = db.QueryRow("SELECT nextval('user_id_seq')").Scan(&userID)
	if err != nil {
		return err
	}

	// Insert the new user with the fetched userID
	_, err = db.Exec("INSERT INTO users (user_id, username, password_hash) VALUES (?, ?, ?)", userID, username, hashedPassword)
	if err != nil {
		if err.Error() == "Constraint Error: UNIQUE constraint failed: users.username" {
			return errors.New("username already exists")
		}
		return err
	}

	return nil
}

// AuthenticateUser checks if username and password match any record in the database
func AuthenticateUser(db *sql.DB, username, password string) (bool, error) {
	var storedHash string
	err := db.QueryRow("SELECT password_hash FROM users WHERE username = ?", username).Scan(&storedHash)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil // User not found
		}
		return false, err
	}

	return CheckPasswordHash(password, storedHash), nil
}
