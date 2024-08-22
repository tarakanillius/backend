package handlers

import (
	"encoding/json"
	"my-app/internal/utils"
	"net/http"
)

func GenerateRecommendationsHandler(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		Products []string `json:"products"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		return
	}

	// Generate recommendations using LangChain
	recommendations, err := utils.GenerateRecommendations(requestData.Products)
	if err != nil {
		http.Error(w, "Failed to generate recommendations", http.StatusInternalServerError)
		return
	}

	// Respond with the generated recommendations
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"recommendations": recommendations})
}
