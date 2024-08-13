//search_product_by_id.go
package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/bson"
	"my-app/internal/models"
	"my-app/internal/utils"
	"net/http"
	"strconv"
	"github.com/joho/godotenv"
  "os"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
}

func GetProductByID(w http.ResponseWriter, r *http.Request) {
    productID := chi.URLParam(r, "id")
    if productID == "" {
        http.Error(w, "Product ID is required", http.StatusBadRequest)
        return
    }

    db := utils.GetDB()

    var product models.Product
    err := db.Collection(os.Getenv("MONGODB_COLLECTION_NAME")).FindOne(context.TODO(), bson.M{"_id": productID}).Decode(&product)
    if err != nil {
        http.Error(w, fmt.Sprintf("Product not found: %v", err), http.StatusNotFound)
        return
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

