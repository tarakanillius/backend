package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"my-app/internal/models"
	"my-app/internal/utils"
	"net/http"
	"os"
	"strconv"
)

// GetProductByID handles requests to fetch a product by its ID.
func GetProductByID(w http.ResponseWriter, r *http.Request) {
	productID := r.URL.Query().Get("id")
	if productID == "" {
		http.Error(w, "Product ID is required", http.StatusBadRequest)
		return
	}

	db := utils.GetDB()

	var product models.Product
	err := db.Collection(os.Getenv("MONGODB_COLLECTION")).FindOne(context.TODO(), bson.M{"_id": productID}).Decode(&product)
	if err != nil {
		http.Error(w, fmt.Sprintf("Product not found: %v", err), http.StatusNotFound)
		return
	}

	// Check if the product already has a nutrition score and grade
	if product.NutritionScore == 0 && product.NutriScoreGrade == "unknown" {
		// Calculate Nutrition Score
		nutritionScore, nutriScoreGrade := utils.CalculateNutritionScore(
			product.Nutriments.EnergyKj100g,
			product.Nutriments.Sugars100g,
			product.Nutriments.SaturatedFat100g,
			product.Nutriments.Salt100g,
			product.Nutriments.Carbohydrates100g,
			product.Nutriments.Fiber100g,
			product.Nutriments.Proteins100g,
		)
		product.NutritionScore = nutritionScore
		product.NutriScoreGrade = nutriScoreGrade
	}

	// Set image URL
	maximg, _ := strconv.Atoi(product.MaxImgID)
	if maximg > 0 {
		product.ImageURL = utils.ComputeImageURL(productID)
	} else {
		product.ImageURL = ""
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(product); err != nil {
		http.Error(w, fmt.Sprintf("Failed to encode product: %v", err), http.StatusInternalServerError)
	}
}
