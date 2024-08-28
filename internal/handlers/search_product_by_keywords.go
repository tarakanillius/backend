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

	findOptions := options.Find().SetLimit(4)
	cur, err := db.Collection(os.Getenv("MONGODB_COLLECTION")).Find(context.TODO(), filter, findOptions)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to search products: %v", err), http.StatusInternalServerError)
		return
	}
	defer cur.Close(context.TODO())

	productScores := make(map[string]int)
	for cur.Next(context.TODO()) {
		var product models.Product
		if err := cur.Decode(&product); err != nil {
			http.Error(w, fmt.Sprintf("Failed to decode product: %v", err), http.StatusInternalServerError)
			return
		}

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

	var sortedProducts []models.Product
	for id := range productScores {
		var product models.Product
		err := db.Collection(os.Getenv("MONGODB_COLLECTION")).FindOne(context.TODO(), bson.M{"_id": id}).Decode(&product)
		if err != nil {
			continue
		}

			if product.NutritionGrade == "unknown" {
		fmt.Printf("Product data: %+v\n", product) // Debug print

		// Convert map[string]interface{} to map[string]float64
		nutriData := make(map[string]float64)
		for k, v := range product.Nutriscore["2023"].Data {
			switch val := v.(type) {
			case float64:
				nutriData[k] = val
			case int32:
				nutriData[k] = float64(val)
			case map[string]interface{}:
				for mk, mv := range val {
					if mvFloat, ok := mv.(float64); ok {
						nutriData[mk] = mvFloat
					}
				}
			case []interface{}:
				for _, item := range val {
					if itemMap, ok := item.(map[string]interface{}); ok {
						for mk, mv := range itemMap {
							if mvFloat, ok := mv.(float64); ok {
								nutriData[mk] = mvFloat
							}
						}
					}
				}
			default:
				fmt.Printf("Unexpected type for key %s: %T\n", k, v)
			}
		}

		fmt.Printf("NutriData: %+v\n", nutriData) // Debug print

		// Determine product type based on both 2021 and 2023 data
		productType := utils.DetermineProductType(product.Nutriscore["2021"].Data, product.Nutriscore["2023"].Data)
		fmt.Printf("Product Type: %s\n", productType) // Debug print

		// Convert float64 values to the required format for functions
		energy := int(nutriData["energy"])
		sugar := nutriData["sugar"]
		saturates := nutriData["saturates"]
		sodium := int(nutriData["sodium"])
		fiber := nutriData["fiber"]
		protein := nutriData["protein"]
		fruitVegNutSeed := nutriData["fruit_vegetable_legume_nut_seed"]
		energyFromSFA := float64(nutriData["energy_from_sfa"])
		totalFat := nutriData["total_fat"]
		salt := nutriData["salt"]

		var score int
		switch productType {
		case "beverage":
			score = utils.CalculateBeverageScore(energy, sugar, saturates, sodium, fruitVegNutSeed, fiber, protein, product.IngredientsNonNutritiveSweeteners > 0)
		case "fat_oil_nut_seed":
			score = utils.CalculateFatsOilsNutsSeedsScore(energy, totalFat, saturates, sodium, sugar, energyFromSFA, energyFromSFA, fruitVegNutSeed, fiber, protein)
		case "red_meat":
			score = utils.CalculateRedMeatScore(energy, sugar, saturates, sodium, salt, protein)
		case "cheese":
			score = utils.CalculateCheeseScore(energy, sugar, saturates, sodium, protein)
		default:
			score = utils.CalculateGeneralFoodScore(energy, sugar, saturates, sodium, fruitVegNutSeed, fiber, protein)
		}
		product.NutritionScore = score
	}
		sortedProducts = append(sortedProducts, product)
	}

	sort.Slice(sortedProducts, func(i, j int) bool {
		return sortedProducts[i].NutritionScore > sortedProducts[j].NutritionScore
	})

	for i, product := range sortedProducts {
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
