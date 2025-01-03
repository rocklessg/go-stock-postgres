 package main

 import(
	"fmt"
	"log"
	"net/http" // used to access the request and response object of the api

	"go-stock-api/router"
 )

 func main() {
	//  Define the port number
	r := router.Router()
	fmt.Println("Starting server on the port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
	
 }