package database

import (
	"database/sql"
	"fmt"
	"log"
	"os" // used to read the environment variable

	_ "github.com/lib/pq" // postgres golang driver
	"github.com/joho/godotenv" // package used to read the .env file
)

// StockDbContext initializes a connection to the PostgreSQL database using
// the connection string from the .env file. It returns a pointer to the
// sql.DB instance and an error if any occurs.
func StockDbContext() (*sql.DB, error) {

	// Load environment variables from .env file
	fmt.Println("Loading environment variables...")

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Get the PostgreSQL connection string from the environment variable
	fmt.Println("Reading connection string from .env file...")
	connStr := os.Getenv("POSTGRES_CONNECTION_STRING")
	if connStr == "" {
		return nil, fmt.Errorf("POSTGRES_CONNECTION_STRING is not set in the .env file")
	}

	// Open a connection to the database
	fmt.Println("Connecting to the database...")
	
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
	log.Println("Connected to the database successfully.")
	return db, nil
}