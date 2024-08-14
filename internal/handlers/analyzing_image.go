package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"my-app/internal/utils"
	"net/http"
)

func AnalyzeImageHandler(w http.ResponseWriter, r *http.Request) {
	// Limit the size of the file to avoid large uploads
	const maxFileSize = 10 * 1024 * 1024 // 10 MB

	// Retrieve the userID from the form data
	userID := r.FormValue("userID")
	if userID == "" {
		http.Error(w, "Missing userID in request", http.StatusBadRequest)
		return
	}

	// Read the file from the request body
	file, _, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Failed to retrieve file from request", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Check the size of the file
	buf := new(bytes.Buffer)
	if _, err := io.CopyN(buf, file, maxFileSize+1); err != nil && err != io.EOF {
		http.Error(w, "Failed to read file", http.StatusInternalServerError)
		return
	}
	if buf.Len() > maxFileSize {
		http.Error(w, "File is too large", http.StatusRequestEntityTooLarge)
		return
	}

	// Send the image data to the OpenAI API
	imageData := buf.Bytes()
	keywords, err := utils.AnalyzeImage(imageData)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to analyze image: %v", err), http.StatusInternalServerError)
		return
	}

	// Respond with the extracted keywords and userID
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{
		"keywords": keywords,
		"userID":   userID,
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, fmt.Sprintf("Failed to encode response: %v", err), http.StatusInternalServerError)
	}
}
