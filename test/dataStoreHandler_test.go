package test

import (
	"context"
	"math/big"
	"os"
	"testing"

	"go-stock-api/database"
	"go-stock-api/middleware"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func setupTestDB(t *testing.T) *pgxpool.Pool {
// Load environment variables from .env file

err := godotenv.Load()
if err != nil {
	t.Fatalf("Error loading .env file: %v", err)
}

	// Reading connection string from .env file...")
	connString := os.Getenv("POSTGRES_CONNECTION_STRING")
	if connString == "" {
		t.Fatalf("Unable to read CONNECTION_STRING from the .env file")
	}

	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		t.Fatalf("Unable to parse connection string: %v", err)
	}

	dbpool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		t.Fatalf("Unable to create connection pool: %v", err)
	}

	return dbpool
}

func TestAddStock(t *testing.T) {
    db := setupTestDB(t)
    defer db.Close()

    stock := database.CreateStockParams{
        Name: "Test Stock1",
        Price: pgtype.Numeric{
            Int:   big.NewInt(10000), // Represents 100.00
            Exp:   -2,               // Exponent for 2 decimal places
            Valid: true,             // Mark the value as valid
        },
        Company: "Test Company",
    }

    id := middleware.AddStock(stock, db)
    assert.Equal(t, int32(10), id)
}

func TestGetAllStocks(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	stocks, err := middleware.GetAllStocks(db)
	assert.NoError(t, err)
	assert.Len(t, stocks, 2)
}

func TestGetStockById(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	stock, err := middleware.GetStockById(9, db)
	assert.NoError(t, err)
	assert.Equal(t, "Test Stock", stock.Name)
	assert.Equal(t, pgtype.Numeric{Int: big.NewInt(10000), Exp: -2, Valid: true}, stock.Price)
}

func TestEditStock(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	stock := database.UpdateStockParams{
		Name:    "Colbat",
		Price:   pgtype.Numeric{Int: big.NewInt(15000), Exp: -2, Valid: true},
		Company: "Combat Inc",
	}

	message := middleware.EditStock(9, stock, db)
	assert.Equal(t, "Stock updated successfully", message)

	updatedStock, err := middleware.GetStockById(9, db)
	assert.NoError(t, err)
	assert.Equal(t, "Colbat", updatedStock.Name)
	assert.Equal(t, pgtype.Numeric{Int: big.NewInt(15000), Exp: -2, Valid: true}, updatedStock.Price)
}

func TestRemoveStock(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	message, err := middleware.RemoveStock(11, db)
	assert.NoError(t, err)
	assert.Equal(t, "Stock deleted successfully", message)

	_, err = middleware.GetStockById(11, db)
	assert.Error(t, err)
}
