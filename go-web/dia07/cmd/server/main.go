package main

import (
	"fmt"
	"net/http"

	"github.com/gabztrivelato/bgw7/dia07/cmd/server/controllers"
	"github.com/gabztrivelato/bgw7/dia07/internal/domain"
	"github.com/gabztrivelato/bgw7/dia07/internal/product"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	db := make(map[int]domain.Product)
	repository := product.NewRepository(db)
	repository.LoadData()
	controller := controllers.NewProductController(repository)

	r.Route("/products", func(r chi.Router) {
		r.Get("/{id}", controller.GetById)
		r.Post("/", controller.Create)
		r.Put("/{id}", controller.Update)
		r.Patch("/{id}", controller.Patch)
		r.Delete("/{id}", controller.Delete)
	})

	fmt.Println("Server running at port 8080")
	http.ListenAndServe(":8080", r)
}
