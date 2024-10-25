# Go Authentication System with DuckDB

A simple and secure authentication system built with Go and DuckDB, featuring user registration, login, and role management capabilities.

## Features

- User registration with secure password hashing using bcrypt
- User authentication (login)
- Role-based access control infrastructure
- User listing endpoint
- Persistent storage using DuckDB
- RESTful API endpoints

## Prerequisites

- Go 1.x or higher
- DuckDB

## Installation

1. Clone the repository
2. Install the required dependencies:
```bash
go get github.com/marcboeker/go-duckdb
go get golang.org/x/crypto/bcrypt
```

## Project Structure

```
.
├── admin.go       # Admin-related functionality and user management
├── auth.go        # Core authentication logic and password handling
├── db.go          # Database initialization and setup
├── handlers.go    # HTTP request handlers
├── main.go        # Application entry point
└── README.md      # This file
```

## Database Schema

The system uses three main tables:

- `users`: Stores user information
  - `user_id`: Integer (Primary Key)
  - `username`: VARCHAR (Unique)
  - `password_hash`: VARCHAR

- `roles`: Stores available roles
  - `role_id`: Integer (Primary Key)
  - `role_name`: VARCHAR (Unique)

- `user_roles`: Maps users to roles
  - `user_id`: Integer (Foreign Key)
  - `role_id`: Integer (Foreign Key)

## API Endpoints

### Register User
```
POST /register
Content-Type: application/json

{
    "username": "string",
    "password": "string"
}
```

### Login
```
POST /login
Content-Type: application/json

{
    "username": "string",
    "password": "string"
}
```

### List Users
```
GET /users
```

## Security Features

- Password hashing using bcrypt
- No password storage in plain text
- Unique username constraints
- Password hash excluded from user listings
- Input validation for registration and login

## Running the Application

1. Start the server:
```bash
go run .
```

2. The server will start on port 8080 by default
3. The database file (auth.db) will be created automatically if it doesn't exist

## Usage Example

Register a new user:
```bash
curl -X POST http://localhost:8080/register \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser","password":"secretpassword"}'
```

Login:
```bash
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser","password":"secretpassword"}'
```

List users:
```bash
curl http://localhost:8080/users
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin feature/my-new-feature`)
5. Create a new Pull Request

## Future Improvements

- Add JWT token-based authentication
- Implement password reset functionality
- Add user profile management
- Implement rate limiting
- Add request logging
- Add input validation middleware
- Implement session management
- Add API documentation using Swagger/OpenAPI
