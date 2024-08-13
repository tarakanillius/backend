package utils

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

var db *mongo.Database

func ConnectDB() error {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("error loading .env file: %v", err)
	}

	// Get MongoDB details from environment variables
	mongoURI := os.Getenv("MONGODB_URI")
	mongoUser := os.Getenv("MONGODB_LOGIN")
	mongoPassword := os.Getenv("MONGODB_PASSWORD")
	mongoDatabase := os.Getenv("MONGODB_DATABASE")

	if mongoURI == "" || mongoUser == "" || mongoPassword == "" || mongoDatabase == "" {
		return fmt.Errorf("MongoDB environment variables are not set correctly")
	}

	// Setup client options with Auth
	clientOptions := options.Client().ApplyURI(mongoURI).SetAuth(options.Credential{
		Username: mongoUser,
		Password: mongoPassword,
	})

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return fmt.Errorf("failed to connect to MongoDB: %v", err)
	}

	// Ping MongoDB to test connection
	if err := client.Ping(context.Background(), nil); err != nil {
		return fmt.Errorf("failed to ping MongoDB: %v", err)
	}

	// Set the database
	db = client.Database(mongoDatabase)
	return nil
}

func GetDB() *mongo.Database {
	return db
}
