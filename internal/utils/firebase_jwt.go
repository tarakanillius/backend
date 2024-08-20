package utils

import (
	"context"
	"firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"fmt"
	"google.golang.org/api/option"
	"log"
	"net/http"
	"os"
	"strings"
)

var firebaseAuth *auth.Client

// InitFirebase initializes the Firebase Admin SDK
var app *firebase.App

func InitFirebase() {
	// Initialize Firebase
	opt := option.WithCredentialsJSON([]byte(os.Getenv("FIREBASE_CONFIG")))
	var err error
	app, err = firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v", err)
	}
}

func GetFirebaseApp() *firebase.App {
	return app
}

// VerifyToken verifies the JWT token with Firebase
func VerifyToken(idToken string) (*auth.Token, error) {
	// Verify the ID token and check if it's valid
	token, err := firebaseAuth.VerifyIDToken(context.Background(), idToken)
	if err != nil {
		return nil, fmt.Errorf("error verifying ID token: %v", err)
	}
	return token, nil
}

func FirebaseAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract the token from the header
		authorizationHeader := r.Header.Get("Authorization")
		if authorizationHeader == "" {
			http.Error(w, "Authorization header is required", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimSpace(strings.TrimPrefix(authorizationHeader, "Bearer"))
		if tokenString == "" {
			http.Error(w, "Authorization token is required", http.StatusUnauthorized)
			return
		}

		// Initialize Firebase App
		ctx := context.Background()
		opt := option.WithCredentialsFile("serviceAccountKey.json")
		app, err := firebase.NewApp(ctx, nil, opt)
		if err != nil {
			log.Fatalf("Failed to initialize Firebase app: %v", err)
			http.Error(w, "Failed to initialize Firebase app", http.StatusInternalServerError)
			return
		}

		// Initialize Auth Client
		client, err := app.Auth(ctx)
		if err != nil {
			log.Fatalf("Failed to create Firebase auth client: %v", err)
			http.Error(w, "Failed to create Firebase auth client", http.StatusInternalServerError)
			return
		}

		// Verify the token
		token, err := client.VerifyIDToken(ctx, tokenString)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Attach user ID to context (optional)
		ctx = context.WithValue(r.Context(), "user", token.UID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
