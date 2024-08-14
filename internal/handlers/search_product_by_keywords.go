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
	"sort"
	"strconv"
	"strings"
)

func GetProductByKeywords(w http.ResponseWriter, r *http.Request) {
	keywordStr := r.URL.Query().Get("keywords")
	if keywordStr == "" {
		http.Error(w, "Keywords query parameter is required", http.StatusBadRequest)
		return
	}

	keywords := strings.FieldsFunc(keywordStr, func(c rune) bool {
		return c == ','
	})

	if len(keywords) == 0 {
		http.Error(w, "No valid keywords provided", http.StatusBadRequest)
		return
	}

	db := utils.GetDB()
	filter := bson.M{"_keywords": bson.M{"$in": keywords}}

	// Find matching products
	findOptions := options.Find().SetLimit(4) // Adjust limit as needed
	cur, err := db.Collection(os.Getenv("MONGODB_COLLECTION")).Find(context.TODO(), filter, findOptions)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to search products: %v", err), http.StatusInternalServerError)
		return
	}
	defer cur.Close(context.TODO())

	// Create a map to store product matches and their scores
	productScores := make(map[string]int)
	for cur.Next(context.TODO()) {
		var product models.Product
		if err := cur.Decode(&product); err != nil {
			http.Error(w, fmt.Sprintf("Failed to decode product: %v", err), http.StatusInternalServerError)
			return
		}

		// Calculate the match score
		score := 0
		for _, keyword := range keywords {
			for _, productKeyword := range product.Keywords {
				if strings.Contains(strings.ToLower(productKeyword), strings.ToLower(keyword)) {
					score++
					break
				}
			}
		}
		productScores[product.ID] = score
	}

	if err := cur.Err(); err != nil {
		http.Error(w, fmt.Sprintf("Cursor error: %v", err), http.StatusInternalServerError)
		return
	}

	// Sort products by their score
	var sortedProducts []models.Product
	for id, score := range productScores {
		var product models.Product
		err := db.Collection(os.Getenv("MONGODB_COLLECTION")).FindOne(context.TODO(), bson.M{"_id": id}).Decode(&product)
		if err != nil {
			continue
		}
		product.NutritionScore = score // Adding the score to the product
		sortedProducts = append(sortedProducts, product)
	}

	// Sort products by their score in descending order
	sort.Slice(sortedProducts, func(i, j int) bool {
		return sortedProducts[i].NutritionScore > sortedProducts[j].NutritionScore
	})

	// Update image URLs and calculate nutrition scores
	for i, product := range sortedProducts {
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
			sortedProducts[i].NutritionScore = nutritionScore
			sortedProducts[i].NutriScoreGrade = nutriScoreGrade
		}

		// Set image URL
		maximg, _ := strconv.Atoi(product.MaxImgID)
		if maximg > 0 {
			sortedProducts[i].ImageURL = utils.ComputeImageURL(product.ID)
		} else {
			sortedProducts[i].ImageURL = ""
		}
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(sortedProducts); err != nil {
		http.Error(w, fmt.Sprintf("Failed to encode products: %v", err), http.StatusInternalServerError)
	}
}
