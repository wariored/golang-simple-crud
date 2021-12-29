package routes

import (
	"context"
	"github.com/go-chi/chi/v5"
	"net/http"
	"supplier/handlers"
	"supplier/models"
)

func RegisterRoutes(r *chi.Mux) {
	r.Route("/suppliers", func(r chi.Router) {
		r.Post("/", handlers.CreateSupplier)
		r.Route("/{supplierID}", func(r chi.Router) {
			r.Use(supplierCtx)
			r.Get("/", handlers.GetSupplier)
			r.Put("/", handlers.UpdateSupplier)
			r.Delete("/", handlers.DeleteSupplier)
		})
	})
}

func supplierCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		supplierID := chi.URLParam(r, "supplierID")
		supplier, err := models.GetSupplierByID(supplierID)
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}
		ctx := context.WithValue(r.Context(), "supplier", supplier)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
