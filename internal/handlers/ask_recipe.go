package handlers

import (
	"encoding/json"
	"my-app/internal/utils"
	"net/http"
)

func GenerateReceiptHandler(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		Products []string `json:"products"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		return
	}

	// Generate receipt using LangChain
	receipt, err := utils.GenerateReceipt(requestData.Products)
	if err != nil {
		http.Error(w, "Failed to generate receipt", http.StatusInternalServerError)
		return
	}

	// Respond with the generated receipt
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"receipt": receipt})
}
