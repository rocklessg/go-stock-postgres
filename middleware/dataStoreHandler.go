package middleware

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"go-stock-api/database"
	"go-stock-api/models"

	//_ "github.com/lib/pq"
)

func AddStock(stock database.CreateStockParams, db *sql.DB) int64 {
    // Initialize the sqlc queries
    queries := database.New(db)

    // Vall the CreateStock method from the sqlc generated code
    newStock, err := queries.CreateStock(context.Background(), stock) 
    if err != nil {
        log.Fatalf("Unable to execute the query. %v", err)
    }

    fmt.Printf("Last inserted stock ID: %v\n", newStock.ID)
    return newStock.ID
}

func GetAllStocks(db *sql.DB) ([]database.Stock, error) {
  
    queries := database.New(db)

    // Call the ListStocks method from the sqlc generated code
    stocks, err := queries.ListStocks(context.Background())
    if err != nil {
        return nil, fmt.Errorf("unable to fetch stocks: %v", err) 
    }
    
    return stocks, nil
}

func GetStockById(id int64, db *sql.DB) (database.Stock, error) {
    
    queries := database.New(db)

    stock, err := queries.GetStock(context.Background(), id)
    if err != nil {
        return database.Stock{}, fmt.Errorf("unable to get stock by ID: %v", err)
    }
    return stock, nil
}

func EditStock(id int64, stock database.UpdateStockParams, db *sql.DB) string {
  
    queries := database.New(db)

    
    // Call the UpdateStock method from the sqlc generated code
    err := queries.UpdateStock(context.Background(), stock)
    if err != nil {
        log.Fatalf("Unable to execute the update query. %v", err)
    }  

    fmt.Printf("stock updated successfully %v", id)
    return "Stock updated successfully"
}

func RemoveStock(id int64, db *sql.DB) int64 {
    
    query := `DELETE FROM stocks WHERE id = $1`
    result, err := db.Exec(query, id)

    if err != nil {
        log.Fatalf("Unable to execute delete query. %v", err)
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        log.Fatalf("Unable to fetch rows affected. %v", err)
    }

    fmt.Printf("Total rows/record affected %v", rowsAffected)
    return rowsAffected
}