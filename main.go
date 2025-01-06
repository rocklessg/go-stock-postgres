package main

import (
	"fmt"
	"log"
	"net/http" // used to access the request and response object of the api

	"go-stock-api/database"
	"go-stock-api/router"
)

 func main() {

	// Initialize the database on application startup
	db, err := database.StockDbContext()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	//  Define the port number
	r := router.Router(db)
	fmt.Println("Server Started and running on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", r))
	
 }