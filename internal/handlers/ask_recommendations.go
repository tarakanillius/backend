//askRecommendations.go
package handlers

import (
	"encoding/json"
	"net/http"
	"my-app/internal/utils"
)


// GenerateRecommendationsHandler handles the recommendations request.
func GenerateRecommendationsHandler(w http.ResponseWriter, r *http.Request) {
    var requestData struct {
        Products []string `json:"products"`
    }

    if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    recommendations, err := utils.GenerateRecommendations(requestData.Products)
    if err != nil {
        http.Error(w, "Failed to generate recommendations", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(recommendations); err != nil {
        http.Error(w, "Failed to encode response", http.StatusInternalServerError)
    }
}
