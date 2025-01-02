package database

import (
	"database/sql"
	"fmt"
	"log"
	"os" // used to read the environment variable

	_ "github.com/lib/pq" // package used to read the .env file
	"github.com/joho/godotenv" // postgres golang driver
)

// StockDbContext initializes a connection to the PostgreSQL database using
// the connection string from the .env file. It returns a pointer to the
// sql.DB instance and an error if any occurs.
func StockDbContext() (*sql.DB, error) {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Get the PostgreSQL connection string from the environment variable
	connStr := os.Getenv("POSTGRES_CONNECTION_STRING")
	if connStr == "" {
		log.Fatalf("POSTGRES_CONNECTION_STRING not set in .env file")
	}

	// Open a connection to the database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}

	// Verify the connection to the database
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("error connecting to the database: %v", err)
	}

	// Return the database connection
	return db, nil
}