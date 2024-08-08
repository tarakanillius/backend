package routes

import (
	"github.com/go-chi/chi/v5"
	"my-app/handlers"
)

// SetupRoutes sets up the routes for the application
func SetupRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Post("/analyze", handlers.AnalyzeImageHandler)
	r.Get("/product/{id}", handlers.GetProductByID)
	r.Get("/search", handlers.GetProductByKeywords)
	r.Post("/receipt", handlers.GenerateReceiptHandler)
	r.Post("/recommendations", handlers.GenerateRecommendationsHandler)

	return r
}
