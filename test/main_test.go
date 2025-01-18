package test

import (
	"testing"
)

func TestMain(m *testing.M) {
	// Initialize database
	db = setupTestDB()
	defer db.Close()

	// Run tests
	m.Run()

	// Clean up database
	teardownTestDB()
}