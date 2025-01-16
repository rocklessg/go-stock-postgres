package database

import (
	"context"
	"fmt"
	"log"
	"os" // used to read the environment variable

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv" // package used to read the .env file
)

// StockDbContext initializes a connection to the PostgreSQL database using
// the connection string from the .env file. It returns a pointer to the
// sql.DB instance and an error if any occurs.
func StockDbContext() (*pgxpool.Pool, error) {

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

	// Create a new connection pool
	fmt.Println("Connecting to the database...")
	
	pool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		return nil, fmt.Errorf("unable to create database connection pool: %v", err)
	}

	// Verify the connection to the database
	err = pool.Ping(context.Background())
	if err != nil {
		return nil, fmt.Errorf("error connecting to the database: %v", err)
	}	

	// Return the database connection
	log.Println("Connected to the database successfully.")
	return pool, nil
}