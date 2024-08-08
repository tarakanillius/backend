package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"my-app/models"
	"my-app/utils"
	"net/http"
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
            product.ImageURL = utils.ComputeImageURL(product.ID)
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
