package router

import (
	"database/sql"
	controller "go-stock-api/controllers"
	"net/http"

	"github.com/gorilla/mux" // used to get the params from the route
)

func Router(db *sql.DB) *mux.Router {

	router := mux.NewRouter()	

	router.HandleFunc("/api/stock/{id}", func(w http.ResponseWriter, r *http.Request) {
		controller.GetStock(w, r, db)}).Methods("GET", "OPTIONS")

	router.HandleFunc("/api/stock", func(w http.ResponseWriter, r *http.Request) {
		controller.GetStocks(w, r, db)}).Methods("GET")

	router.HandleFunc("/api/stock/add", func(w http.ResponseWriter, r *http.Request) {
		controller.CreateStock(w, r, db)}).Methods("POST")

	router.HandleFunc("/api/stock/update/{id}", func(w http.ResponseWriter, r *http.Request) {
		controller.UpdateStock(w, r, db)}).Methods("PUT")

	router.HandleFunc("/api/stock/delete/{id}", func(w http.ResponseWriter, r *http.Request) {
		controller.DeleteStock(w, r, db)}).Methods("DELETE")

	return router
}