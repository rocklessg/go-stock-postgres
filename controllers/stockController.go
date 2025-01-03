package controllers

import (
	"encoding/json" // package to encode and decode the json into struct and vice versa
	"fmt"
	"net/http"
	"strconv" // package used to covert string to int

	"go-stock-api/middleware"
	"go-stock-api/models"

	"github.com/gorilla/mux"
)

type response struct {
	ID		 int64  `json:"id"`
	Message  string `json:"message"`
}

func CreateStock(w http.ResponseWriter, r *http.Request) {
	
	var stock models.Stock

	// Decode the incoming Stock json to the Stock struct
	err := json.NewDecoder(r.Body).Decode(&stock)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	insertStockResult := middleware.AddStock(stock)

	// Format and Return the response
	res := response{
		ID: insertStockResult,
		Message: "Stock created successfully",
	}
	json.NewEncoder(w).Encode(res)
}

func GetStocks(w http.ResponseWriter, r *http.Request) {

	stocks, err := middleware.GetAllStocks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Format and Return the response
	json.NewEncoder(w).Encode(stocks)
}

func GetStock(w http.ResponseWriter, r *http.Request) {
	// Get the stockId from the request params, key is "id"
	var params = mux.Vars(r)

	// Convert the id type from string to int
	 id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Get the stock by id
	stock, err := middleware.GetStockById(int64(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Format and Return the response
	json.NewEncoder(w).Encode(stock)
}

func UpdateStock(w http.ResponseWriter, r *http.Request) {

	// Get the stockId from the request params, key is "id"
	var params = mux.Vars(r)

	// Convert the id type from string to int
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var stock models.Stock

	// Decode the incoming Stock json to the Stock struct
	err = json.NewDecoder(r.Body).Decode(&stock)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Update the stock
	updateStockResult := middleware.EditStock(int64(id), stock)

	// format the response message
	msg := fmt.Sprintf("Stock updated successfully. Total rows/record affected %v", updateStockResult)

	// Format and Return the response
	res := response{
		ID: updateStockResult,
		Message: msg,
	}
	json.NewEncoder(w).Encode(res)
}

func DeleteStock(w http.ResponseWriter, r *http.Request) {
	// Get the stockId from the request params, key is "id"
	var params = mux.Vars(r)

	// Convert the id type from string to int
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Delete the stock
	deleteStockResult := middleware.RemoveStock(int64(id))

	// format the response message
	msg := fmt.Sprintf("Stock deleted successfully. Total rows/record affected %v", deleteStockResult)

	// Format and Return the response
	res := response{
		ID: deleteStockResult,
		Message: msg,
	}
	json.NewEncoder(w).Encode(res)
}