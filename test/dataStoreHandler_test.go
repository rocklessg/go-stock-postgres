// package test

// import (
// 	"context"
// 	"math/big"
// 	"os"
// 	"testing"

// 	"go-stock-api/database"
// 	"go-stock-api/middleware"

// 	"github.com/jackc/pgx/v5/pgtype"
// 	"github.com/jackc/pgx/v5/pgxpool"
// 	"github.com/joho/godotenv"
// 	"github.com/stretchr/testify/assert"
// )

// func setupTestDB(t *testing.T) *pgxpool.Pool {
// // Load environment variables from .env file

// err := godotenv.Load()
// if err != nil {
// 	t.Fatalf("Error loading .env file: %v", err)
// }

// 	// Reading connection string from .env file...")
// 	connString := os.Getenv("POSTGRES_CONNECTION_STRING")
// 	if connString == "" {
// 		t.Fatalf("Unable to read CONNECTION_STRING from the .env file")
// 	}

// 	config, err := pgxpool.ParseConfig(connString)
// 	if err != nil {
// 		t.Fatalf("Unable to parse connection string: %v", err)
// 	}

// 	dbpool, err := pgxpool.NewWithConfig(context.Background(), config)
// 	if err != nil {
// 		t.Fatalf("Unable to create connection pool: %v", err)
// 	}

// 	return dbpool
// }

// func TestAddStock(t *testing.T) {
//     db := setupTestDB(t)
//     defer db.Close()

//     stock := database.CreateStockParams{
//         Name: "Test Stock1",
//         Price: pgtype.Numeric{
//             Int:   big.NewInt(10000), // Represents 100.00
//             Exp:   -2,               // Exponent for 2 decimal places
//             Valid: true,             // Mark the value as valid
//         },
//         Company: "Test Company",
//     }

//     id := middleware.AddStock(stock, db)
//     assert.Equal(t, int32(10), id)
// }

// func TestGetAllStocks(t *testing.T) {
// 	db := setupTestDB(t)
// 	defer db.Close()

// 	stocks, err := middleware.GetAllStocks(db)
// 	assert.NoError(t, err)
// 	assert.Len(t, stocks, 2)
// }

// func TestGetStockById(t *testing.T) {
// 	db := setupTestDB(t)
// 	defer db.Close()

// 	stock, err := middleware.GetStockById(9, db)
// 	assert.NoError(t, err)
// 	assert.Equal(t, "Test Stock", stock.Name)
// 	assert.Equal(t, pgtype.Numeric{Int: big.NewInt(10000), Exp: -2, Valid: true}, stock.Price)
// }

// func TestEditStock(t *testing.T) {
// 	db := setupTestDB(t)
// 	defer db.Close()

// 	stock := database.UpdateStockParams{
// 		Name:    "Colbat",
// 		Price:   pgtype.Numeric{Int: big.NewInt(15000), Exp: -2, Valid: true},
// 		Company: "Combat Inc",
// 	}

// 	message := middleware.EditStock(9, stock, db)
// 	assert.Equal(t, "Stock updated successfully", message)

// 	updatedStock, err := middleware.GetStockById(9, db)
// 	assert.NoError(t, err)
// 	assert.Equal(t, "Colbat", updatedStock.Name)
// 	assert.Equal(t, pgtype.Numeric{Int: big.NewInt(15000), Exp: -2, Valid: true}, updatedStock.Price)
// }

// func TestRemoveStock(t *testing.T) {
// 	db := setupTestDB(t)
// 	defer db.Close()

// 	message, err := middleware.RemoveStock(11, db)
// 	assert.NoError(t, err)
// 	assert.Equal(t, "Stock deleted successfully", message)

// 	_, err = middleware.GetStockById(11, db)
// 	assert.Error(t, err)
// }

package test

import (
	//"context"
	"math/big"
	//"os"
	"testing"

	"go-stock-api/database"
	"go-stock-api/middleware"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	//"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

var db *pgxpool.Pool

// Setup function to initialize database connection
// func setupTestDB() *pgxpool.Pool {
// 	err := godotenv.Load()
// 	if err != nil {
// 		panic("Error loading .env file: " + err.Error())
// 	}

// 	// Reading connection string from .env file...")
// 	connString := os.Getenv("POSTGRES_TEST_CONNECTION_STRING")
// 	if connString == "" {
// 		panic("Unable to read CONNECTION_STRING from the .env file")
// 	}

// 	conn, err := pgxpool.New(context.Background(), connString)
// 	if err != nil {
// 		panic("Unable to connect to test database: " + err.Error())
// 	}
// 	return conn
// }

// Teardown function to clean up database after tests
// func teardownTestDB() {
// 	_, err := db.Exec(context.Background(), "TRUNCATE stocks RESTART IDENTITY CASCADE;")
// 	if err != nil {
// 		panic("Failed to truncate test database: " + err.Error())
// 	}
// }

// func TestMain(m *testing.M) {
// 	// Initialize database
// 	db = setupTestDB()
// 	defer db.Close()

// 	// Run tests
// 	m.Run()

// 	// Clean up database
// 	teardownTestDB()
// }

func TestAddStock(t *testing.T) {

	stock := database.CreateStockParams{
		Name: "Test Stock1",
		Price: pgtype.Numeric{
			Int:   big.NewInt(10000), // Represents 100.00
			Exp:   -2,                // Exponent for 2 decimal places
			Valid: true,              // Mark the value as valid
		},
		Company: "Test Company",
	}

	id := middleware.AddStock(stock, db)
	assert.NotZero(t, id, "New stock ID should not be zero")
}

func TestGetAllStocks(t *testing.T) {
	// Add a stock for testing
	stock := database.CreateStockParams{
		Name:  "Test Stock 2",
		Price: pgtype.Numeric{
			Int:   big.NewInt(20075), // Represents 200.75
			Exp:   -2,                
			Valid: true,             
		},
		Company: "Test Company",
	}
	middleware.AddStock(stock, db)

	stocks, err := middleware.GetAllStocks(db)
	assert.NoError(t, err, "Fetching all stocks should not return an error")
	assert.NotEmpty(t, stocks, "Stocks list should not be empty")
}

func TestGetStockById(t *testing.T) {
	// Add a stock and fetch it by ID
	stock := database.CreateStockParams{
		Name:  "Stock By ID",
		Price: pgtype.Numeric{
			Int:   big.NewInt(15025), // Represents 150.25
			Exp:   -2,
			Valid: true,
		},
		Company: "Test Company",
	}
	id := middleware.AddStock(stock, db)

	fetchedStock, err := middleware.GetStockById(id, db)
	assert.NoError(t, err, "Fetching stock by ID should not return an error")
	assert.Equal(t, stock.Name, fetchedStock.Name, "Fetched stock name should match")
	assert.Equal(t, stock.Price, fetchedStock.Price, "Fetched stock price should match")
}

func TestEditStock(t *testing.T) {
	// Add a stock, then update it
	stock := database.CreateStockParams{
		Name:  "Editable Stock",
		Price: pgtype.Numeric{
			Int:   big.NewInt(30050), // Represents 300.50
			Exp:   -2,
			Valid: true,
		},
		Company: "Test Company",
	}
	id := middleware.AddStock(stock, db)

	updateParams := database.UpdateStockParams{
		Name:  "Updated Stock",
		Price: pgtype.Numeric{
			Int:   big.NewInt(35075), // Represents 350.75
			Exp:   -2,
			Valid: true,
		},
		Company: "Updated Company",
	}
	message := middleware.EditStock(id, updateParams, db)
	assert.Equal(t, "Stock updated successfully", message)

	// Verify update
	fetchedStock, _ := middleware.GetStockById(id, db)
	assert.Equal(t, updateParams.Name, fetchedStock.Name, "Stock name should be updated")
	assert.Equal(t, updateParams.Price, fetchedStock.Price, "Stock price should be updated")
	assert.Equal(t, updateParams.Company, fetchedStock.Company, "Stock company should be updated")
}

func TestRemoveStock(t *testing.T) {
	// Add a stock, then delete it
	stock := database.CreateStockParams{
		Name:  "Deletable Stock",
		Price: pgtype.Numeric{
			Int:   big.NewInt(50025), // Represents 500.25
			Exp:  -2,
			Valid: true,
		},
		Company: "Test Company",
	}
	id := middleware.AddStock(stock, db)

	message, err := middleware.RemoveStock(id, db)
	assert.NoError(t, err, "Deleting stock should not return an error")
	assert.Equal(t, "Stock deleted successfully", message)

	// Verify deletion
	_, err = middleware.GetStockById(id, db)
	assert.Error(t, err, "Fetching deleted stock should return an error")
}
