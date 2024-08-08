package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"my-app/models"
	"my-app/utils"
	"net/http"
	"strconv"
	"strings"
)

// GetProductByID handles fetching a product by its ID.
func GetProductByID(w http.ResponseWriter, r *http.Request) {
	productID := chi.URLParam(r, "id")
	if productID == "" {
		http.Error(w, "Product ID is required", http.StatusBadRequest)
		return
	}

	db := utils.GetDB()

	var product models.Product
	err := db.Collection("products").FindOne(context.TODO(), bson.M{"_id": productID}).Decode(&product)
	if err != nil {
		http.Error(w, fmt.Sprintf("Product not found: %v", err), http.StatusNotFound)
		return
	}

	maximg, _ := strconv.Atoi(product.MaxImgID)
	if maximg > 0 {
		product.ImageURL = computeImageURLs(productID)
	} else {
		product.ImageURL = "" // No image URL if `max_imgid` is not present or <= 0
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(product); err != nil {
		http.Error(w, fmt.Sprintf("Failed to encode product: %v", err), http.StatusInternalServerError)
	}
}

// SearchProducts handles searching for products by keywords.
func SearchProducts(w http.ResponseWriter, r *http.Request) {
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

	cur, err := db.Collection("products").Find(context.TODO(), filter, findOptions)
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

		maximg, _ := strconv.Atoi(product.MaxImgID)
		if maximg > 0 {
			product.ImageURL = computeImageURLs(product.ID)
		} else {
			product.ImageURL = "" // No image URL if `max_imgid` is not present or <= 0
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

// computeImageURLs generates a URL for the product image.
func computeImageURLs(barcode string) string {
	var url string
	if len(barcode) <= 8 {
		url = fmt.Sprintf("https://images.openfoodfacts.org/images/products/%s/%d.jpg", barcode, 1)
	} else {
		prefix := barcode[:9]
		lastPart := barcode[9:]
		url = fmt.Sprintf("https://images.openfoodfacts.org/images/products/%s/%s/%s/%s/%d.jpg",
			prefix[:3],
			prefix[3:6],
			prefix[6:],
			lastPart,
			1)
	}
	return url
}
