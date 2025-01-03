package router

import (
	"github.com/gorilla/mux" // used to get the params from the route
	controller "go-stock-api/controllers"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/stock/{id}", controller.GetStock).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/stock", controller.GetStocks).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/stock/add", controller.CreateStock).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/stock/update/{id}", controller.UpdateStock).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/stock/delete/{id}", controller.DeleteStock).Methods("DELETE", "OPTIONS")

	return router
}