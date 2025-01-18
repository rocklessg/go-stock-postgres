package test

import (
	"context"
	"os"
	"log"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

// Setup function to initialize database connection
func setupTestDB() *pgxpool.Pool {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Read connection strings
	defaultConnString := os.Getenv("POSTGRES_CONNECTION_STRING") // Connects to the default 'postgres' database
	testConnString := os.Getenv("POSTGRES_TEST_CONNECTION_STRING") // Connects to 'go-stock_test_db'

	if defaultConnString == "" || testConnString == "" {
		log.Fatalf("Connection strings are missing in the .env file")
	}

	// Step 1: Connect to the default 'postgres' database
	defaultConn, err := pgx.Connect(context.Background(), defaultConnString)
	if err != nil {
		log.Fatalf("Unable to connect to the default database: %v", err)
	}
	defer defaultConn.Close(context.Background())

	// Step 2: Check if the test database exists, and create it if not
	_, err = defaultConn.Exec(context.Background(), "CREATE DATABASE go_stock_test_db")
	if err != nil && err.Error() != `ERROR: database "go_stock_test_db" already exists (SQLSTATE 42P04)` {
		log.Fatalf("Failed to create test database: %v", err)
	}

	// Step 3: Connect to the test database
	testConn, err := pgxpool.New(context.Background(), testConnString)
	if err != nil {
		log.Fatalf("Unable to connect to test database: %v", err)
	}

	// Step 4: Ensure the required table exists in the test database
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS stocks (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		price NUMERIC(10, 2) NOT NULL,
		company TEXT,
		created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
	)`
	_, err = testConn.Exec(context.Background(), createTableQuery)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}

	log.Println("Test database setup complete")
	return testConn
}

func teardownTestDB() {
	_, err := db.Exec(context.Background(), "TRUNCATE stocks RESTART IDENTITY CASCADE;")
	if err != nil {
		panic("Failed to truncate test database: " + err.Error())
	}
}