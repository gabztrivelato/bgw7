package main

import (
	"fmt"
	"net/http"

	"github.com/gabztrivelato/bgw7/dia07/cmd/server/controllers"
	"github.com/gabztrivelato/bgw7/dia07/internal/product"
	"github.com/go-chi/chi/v5"
)

func main() {
	product.LoadData()
	r := chi.NewRouter()

	repository := product.NewRepository()
	controller := controllers.NewProductController(repository)

	r.Route("/products", func(r chi.Router) {
		r.Get("/{id}", controller.GetById)
		r.Post("/", controller.Create)
	})

	fmt.Println("Server running at port 8080")
	http.ListenAndServe(":8080", r)
}
