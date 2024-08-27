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

	if product.NutritionScore > 0 || product.NutritionGrade == "unknown" {
		// Convert map[string]interface{} to map[string]float64
		nutriData := make(map[string]float64)
		for k, v := range product.Nutriscore["2021"].Data {
			if val, ok := v.(float64); ok {
				nutriData[k] = val
			} else if val, ok := v.(int); ok {
				nutriData[k] = float64(val)
			}
		}

		productType := utils.DetermineProductType(product.Nutriscore["2021"].Data)

		var score int
		switch productType {
		case "beverage":
			score = utils.CalculateBeverageScore(nutriData)
		default:
			score = utils.CalculateGeneralFoodScore(nutriData)
		}
		product.NutritionScore = score
	}
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
