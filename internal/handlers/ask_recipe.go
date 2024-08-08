//askRecipe.go
package handlers

import (
    "encoding/json"
    "net/http"
    "my-app/internal/utils"
)

// GenerateReceiptHandler handles the receipt generation request.
func GenerateReceiptHandler(w http.ResponseWriter, r *http.Request) {
    var requestData struct {
        Products []string `json:"products"`
    }

    if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    recipe, err := utils.GenerateReceipt(requestData.Products)
    if err != nil {
        http.Error(w, "Failed to generate receipt", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(recipe); err != nil {
        http.Error(w, "Failed to encode response", http.StatusInternalServerError)
    }
}


