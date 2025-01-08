package middleware

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"go-stock-api/database"
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

func RemoveStock(id int64, db *sql.DB) (string, error) {
    
    queries := database.New(db)

    // Call the DeleteStock method from the sqlc generated code
    err := queries.DeleteStock(context.Background(), id) 

    if err != nil {
        return "", fmt.Errorf("unable to execute delete query: %v", err)
    }
    
    fmt.Println("Stock deleted successfully")
    return "Stock deleted successfully", nil
}