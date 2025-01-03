package middleware

import (
    "fmt"
    "log"
    "time"

    "go-stock-api/models"
    "go-stock-api/database"

    _ "github.com/lib/pq"
)

func AddStock(stock models.Stock) int64 {
    db, err := database.StockDbContext()
    if err != nil {
        log.Fatalf("Unable to open the database. %v", err)
    }
    defer db.Close()

    var id int64
    createdAt := time.Now()
    updatedAt := time.Now()

    query := `INSERT INTO stocks (name, price, company, createdAt, updatedAt) VALUES ($1, $2, $3, $4, $5) RETURNING id`
    err = db.QueryRow(query, stock.Name, stock.Price, stock.Company, createdAt, updatedAt).Scan(&id)

    if err != nil {
        log.Fatalf("Unable to execute the query. %v", err)
    }

    fmt.Printf("Last inserted stock ID: %v\n", id)
    return id
}

func GetAllStocks() ([]models.Stock, error) {
    db, err := database.StockDbContext()
    if err != nil {
        return nil, err
    }
    defer db.Close()

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

func GetStockById(id int64) (models.Stock, error) {
    db, err := database.StockDbContext()
    if err != nil {
        log.Fatalf("Unable to open the database. %v", err)
    }
    defer db.Close()

    var stock models.Stock
    query := `SELECT * FROM stocks WHERE id = $1`
    err = db.QueryRow(query, id).Scan(&stock.ID, &stock.Name, &stock.Price, &stock.Company, &stock.CreatedAt, &stock.UpdatedAt)
    
    if err != nil {
        log.Fatalf("Unable to execute the query. %v", err)
    }
    return stock, nil
}

func EditStock(id int64, stock models.Stock) int64 {
    db, err := database.StockDbContext()
    if err != nil {
        log.Fatalf("Unable to open the database. %v", err)
    }
    defer db.Close()

    query := `UPDATE stocks SET name = $1, price = $2, company = $3, updatedAt = $4 WHERE id = $5`
    result, err := db.Exec(query, stock.Name, stock.Price, stock.Company, stock.UpdatedAt, id)
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

func RemoveStock(id int64) int64 {
    db, err := database.StockDbContext()
    if err != nil {
        log.Fatalf("Unable to open the database. %v", err)
    }
    defer db.Close()
 
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