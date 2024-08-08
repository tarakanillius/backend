//dbConnect.go
package utils

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database

// ConnectDB connects to the MongoDB database
func ConnectDB() error {
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGODB_URI"))
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return err
	}
	if err := client.Ping(context.Background(), nil); err != nil {
		return err
	}
	db = client.Database(os.Getenv("MONGODB_NAME"))
	return nil
}

// GetDB returns the MongoDB database instance
func GetDB() *mongo.Database {
	return db
}
