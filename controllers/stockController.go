package controllers

import (
	"database/sql"
	"encoding/json" // package to encode and decode the json into struct and vice versa
	"fmt"
	"net/http"
	"strconv" // package used to covert string to int

	"go-stock-api/database"
	"go-stock-api/middleware"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
)

type response struct {
	ID		 int32  `json:"id"`
	Message  string `json:"message"`
}

func CreateStock(w http.ResponseWriter, r *http.Request, db *pgxpool.Pool) {
	
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
	if err != nil {
		fmt.Println(err)
		// if err != sql.ErrNoRows {
							
		// }	
		// No duplicate found, proceed to creation
			
			// Create new stock
			insertStockResult := middleware.AddStock(stock, db)
			// if insertStockResult == 0 {
			// 	http.Error(w, "Error creating stock", http.StatusInternalServerError)
			// }

			// Format and Return the response
			res := response{
				ID: insertStockResult,
				Message: "Stock created successfully",
			}
			json.NewEncoder(w).Encode(res)						
	}
	if existingStock.Name != "" {
		http.Error(w, "Stock with the same name already exists", http.StatusConflict)
		return
	}		
}

func GetStocks(w http.ResponseWriter, r *http.Request, db *pgxpool.Pool) {

	stocks, err := middleware.GetAllStocks(db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Format and Return the response
	json.NewEncoder(w).Encode(stocks)
}

func GetStock(w http.ResponseWriter, r *http.Request, db *pgxpool.Pool) {
	// Get the stockId from the request params, key is "id"
	var params = mux.Vars(r)

	// Convert the id type from string to int
	 id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Get the stock by id
	stock, err := middleware.GetStockById(int32(id), db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Format and Return the response
	json.NewEncoder(w).Encode(stock)
}

func UpdateStock(w http.ResponseWriter, r *http.Request, db *pgxpool.Pool) {

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
	updateStockResult := middleware.EditStock(int32(id), stock, db)

	// Format and Return the response
	res := response{
		ID: int32(id),
		Message: updateStockResult,
	}
	json.NewEncoder(w).Encode(res)
}

func DeleteStock(w http.ResponseWriter, r *http.Request, db *pgxpool.Pool) {
	
	// Get the stockId from the request params, key is "id"
	var params = mux.Vars(r)

	// Convert the id type from string to int
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Delete the stock
	deleteStockResult, err := middleware.RemoveStock(int32(id), db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Format and Return the response
	res := response{
		ID: int32(id),
		Message: deleteStockResult,
	}
	json.NewEncoder(w).Encode(res)
}