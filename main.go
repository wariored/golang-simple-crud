package main

import (
	"net/http"
	"github.com/go-chi/chi/v5"
	"supplier/routes"
)

func main() {
	r := chi.NewRouter()
	routes.RegisterRoutes(r)
	http.ListenAndServe(":3000", r)
}
