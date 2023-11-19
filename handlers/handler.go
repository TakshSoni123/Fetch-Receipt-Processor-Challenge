package handlers

import (
	"encoding/json"
	"net/http"
	"receipt-processor-backend/helper"
	"receipt-processor-backend/models"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
)

func ProcessReceipt(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	//Create new receipt
	var newReceipt models.Receipt
	err := json.NewDecoder(r.Body).Decode(&newReceipt)
	if err != nil {
		http.Error(w, "Error Decoding receipt json : "+err.Error(), http.StatusBadRequest)
		return
	}

	//Create Id for this receipt
	newReceipt.ID = uuid.New().String()

	// Store it in receipts map with Id : receipt
	models.Receipts[newReceipt.ID] = newReceipt

	// create response data
	response := models.ReceiptID{
		Id: newReceipt.ID,
	}
	data, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error Marshaling return data : "+err.Error(), http.StatusInternalServerError)
		return
	}

	// return id with 200 OK
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)

}

func GetPoints(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	receiptId := params.ByName("id")

	receipt, exists := models.Receipts[receiptId]
	if !exists {
		http.Error(w, "Invalid Id : Receipt not found", http.StatusNotFound)
		return
	}

	points, err := helper.CalculatePoints(receipt)
	if err != nil {
		http.Error(w, "Error calculating points : "+err.Error(), http.StatusInternalServerError)
		return
	}

	// create response data
	response := models.ReceiptPoints{
		Points: points,
	}
	data, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error Marshaling response data : "+err.Error(), http.StatusInternalServerError)
		return
	}

	// return id with 200 OK
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)

}
