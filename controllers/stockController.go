package controllers

import (
	"database/sql"
	"encoding/json" // package to encode and decode the json into struct and vice versa
	"net/http"
	"strconv" // package used to covert string to int

	"go-stock-api/database"
	"go-stock-api/middleware"

	"github.com/gorilla/mux"
)

type response struct {
	ID		 int64  `json:"id"`
	Message  string `json:"message"`
}

func CreateStock(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	
	var stock database.CreateStockParams

	// Decode the incoming Stock json to the Stock struct
	err := json.NewDecoder(r.Body).Decode(&stock)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	queries := database.New(db)

	// Check if the stock name already exists
	existingStock, err := queries.GetStockByName(r.Context(), stock.Name)
	if err != nil || existingStock.Name != "" {
		http.Error(w, "Stock with the same name already exists", http.StatusConflict)
		return
	}

	insertStockResult := middleware.AddStock(stock, db)

	// Format and Return the response
	res := response{
		ID: insertStockResult,
		Message: "Stock created successfully",
	}
	json.NewEncoder(w).Encode(res)
}

func GetStocks(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	stocks, err := middleware.GetAllStocks(db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Format and Return the response
	json.NewEncoder(w).Encode(stocks)
}

func GetStock(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// Get the stockId from the request params, key is "id"
	var params = mux.Vars(r)

	// Convert the id type from string to int
	 id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Get the stock by id
	stock, err := middleware.GetStockById(int64(id), db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Format and Return the response
	json.NewEncoder(w).Encode(stock)
}

func UpdateStock(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	// Get the stockId from the request params, key is "id"
	var params = mux.Vars(r)

	// Convert the id type from string to int
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var stock database.UpdateStockParams

	// Decode the incoming Stock json to the Stock struct
	err = json.NewDecoder(r.Body).Decode(&stock)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Update the stock
	updateStockResult := middleware.EditStock(int64(id), stock, db)

	// Format and Return the response
	res := response{
		ID: int64(id),
		Message: updateStockResult,
	}
	json.NewEncoder(w).Encode(res)
}

func DeleteStock(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	
	// Get the stockId from the request params, key is "id"
	var params = mux.Vars(r)

	// Convert the id type from string to int
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Delete the stock
	deleteStockResult, err := middleware.RemoveStock(int64(id), db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Format and Return the response
	res := response{
		ID: int64(id),
		Message: deleteStockResult,
	}
	json.NewEncoder(w).Encode(res)
}