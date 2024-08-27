package routes

import (
	"github.com/go-chi/chi/v5"
	"my-app/internal/handlers"
)

func SetupRoutes() *chi.Mux {
	r := chi.NewRouter()

	// Apply FirebaseAuthMiddleware globally

	r.Post("/analyze", handlers.AnalyzeImageHandler)
	r.Get("/product", handlers.GetProductByID)
	r.Get("/search", handlers.GetProductByKeywords)
	r.Post("/receipt", handlers.GenerateReceiptHandler)
	r.Post("/recommendations", handlers.GenerateRecommendationsHandler)

	return r
}
