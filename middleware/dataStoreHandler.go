package middleware

import (
    "fmt"
    "log"
    "time"
    "database/sql"

    "go-stock-api/models"

    _ "github.com/lib/pq"
)

func AddStock(stock models.Stock, db *sql.DB) int64 {
    
    var id int64
    createdAt := time.Now()
    updatedAt := time.Now()

    query := `INSERT INTO stocks (name, price, company, created_At, updated_At) VALUES ($1, $2, $3, $4, $5) RETURNING id`
    err := db.QueryRow(query, stock.Name, stock.Price, stock.Company, createdAt, updatedAt).Scan(&id)

    if err != nil {
        log.Fatalf("Unable to execute the query. %v", err)
    }

    fmt.Printf("Last inserted stock ID: %v\n", id)
    return id
}

func GetAllStocks(db *sql.DB) ([]models.Stock, error) {
  
    var stocks []models.Stock

    query := `SELECT * FROM stocks`
    rows, err := db.Query(query)

    if err != nil {
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        var stock models.Stock
        err := rows.Scan(&stock.ID, &stock.Name, &stock.Price, &stock.Company, &stock.CreatedAt, &stock.UpdatedAt)
        if err != nil {
            log.Fatal(err)
        }
        stocks = append(stocks, stock)
    }
    return stocks, nil
}

func GetStockById(id int64, db *sql.DB) (models.Stock, error) {
    
    var stock models.Stock
    query := `SELECT * FROM stocks WHERE id = $1`
    err := db.QueryRow(query, id).Scan(&stock.ID, &stock.Name, &stock.Price, &stock.Company, &stock.CreatedAt, &stock.UpdatedAt)
    
    if err != nil {
        return stock, err
    }
    return stock, nil
}

func EditStock(id int64, stock models.Stock, db *sql.DB) int64 {
  
    query := `UPDATE stocks SET name = $1, price = $2, company = $3, updated_At = $4 WHERE id = $5`
    result, err := db.Exec(query, stock.Name, stock.Price, stock.Company, time.Now(), id)
    if err != nil {
        log.Fatalf("Unable to execute the update query. %v", err)
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        log.Fatalf("Unable to fetch rows affected. %v", err)
    }

    fmt.Printf("Total rows/record affected %v", rowsAffected)
    return rowsAffected
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