package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"my-app/internal/models"
	"my-app/internal/utils"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// GetProductByKeywords handles requests to search for products by keywords.
func GetProductByKeywords(w http.ResponseWriter, r *http.Request) {
	keywordStr := r.URL.Query().Get("keywords")
	if keywordStr == "" {
		http.Error(w, "Keywords query parameter is required", http.StatusBadRequest)
		return
	}

	keywords := strings.FieldsFunc(keywordStr, func(c rune) bool {
		return c == ','
	})

	db := utils.GetDB()
	filter := bson.M{"_keywords": bson.M{"$all": keywords}}
	findOptions := options.Find().SetLimit(4)

	cur, err := db.Collection(os.Getenv("MONGODB_COLLECTION")).Find(context.TODO(), filter, findOptions)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to search products: %v", err), http.StatusInternalServerError)
		return
	}
	defer cur.Close(context.TODO())

	var products []models.Product
	for cur.Next(context.TODO()) {
		var product models.Product
		if err := cur.Decode(&product); err != nil {
			http.Error(w, fmt.Sprintf("Failed to decode product: %v", err), http.StatusInternalServerError)
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
			product.ImageURL = utils.ComputeImageURL(product.ID)
		} else {
			product.ImageURL = ""
		}

		products = append(products, product)
	}

	if err := cur.Err(); err != nil {
		http.Error(w, fmt.Sprintf("Cursor error: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(products); err != nil {
		http.Error(w, fmt.Sprintf("Failed to encode products: %v", err), http.StatusInternalServerError)
	}
}
