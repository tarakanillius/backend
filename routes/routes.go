package routes

import (
	"github.com/go-chi/chi/v5"
	"my-app/handlers"
)

// SetupRoutes sets up the routes for the application
func SetupRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Post("/image/analyze", handlers.AnalyzeImageHandler)
	r.Get("/product/{id}", handlers.GetProductByID)
	r.Get("/search", handlers.SearchProducts)
	r.Post("/generate-receipt", handlers.GenerateReceiptHandler)
	r.Post("/generate-recommendations", handlers.GenerateRecommendationsHandler)

	return r
}
