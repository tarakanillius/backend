package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"my-app/internal/utils"
	"net/http"
)

func AnalyzeImageHandler(w http.ResponseWriter, r *http.Request) {
	const maxFileSize = 10 * 1024 * 1024 // 10 MB

	// Parse userID and image file
	userID := r.FormValue("userID")
	if userID == "" {
		http.Error(w, "Missing userID", http.StatusBadRequest)
		return
	}

	file, _, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Failed to retrieve image", http.StatusBadRequest)
		return
	}
	defer file.Close()

	imageData, err := io.ReadAll(io.LimitReader(file, maxFileSize))
	if err != nil {
		http.Error(w, "Failed to read image", http.StatusInternalServerError)
		return
	}

	// Analyze image using LangChain
	keywords, err := utils.AnalyzeImage(imageData)
	if err != nil {
		http.Error(w, fmt.Sprintf("Image analysis failed: %v", err), http.StatusInternalServerError)
		return
	}

	// Respond with the analysis result
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"keywords": keywords,
	})
}
