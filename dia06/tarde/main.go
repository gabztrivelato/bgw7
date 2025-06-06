package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type Product struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	Code        string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

func main() {
	file, err := os.Open("products.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var products []Product
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&products); err != nil {
		panic(err)
	}

	router := chi.NewRouter()
	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	})

	router.Get("/products", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type", "appliation/json")
		json.NewEncoder(w).Encode(products)
	})

	router.Get("/products/{id}", func(w http.ResponseWriter, r *http.Request) {
		for _, product := range products {
			id, err := strconv.Atoi(chi.URLParam(r, "id"))
			if err != nil {
				http.Error(w, "id must be a valid number", http.StatusBadRequest)
				return
			}

			if product.Id == id {
				w.WriteHeader(http.StatusOK)
				w.Header().Add("Content-Type", "appliation/json")
				json.NewEncoder(w).Encode(product)
			}
		}
	})

	router.Get("/products/search", func(w http.ResponseWriter, r *http.Request) {
		price := r.URL.Query().Get("priceGt")
		if price == "" {
			http.Error(w, "priceGt parameter is required", http.StatusBadRequest)
			return
		}

		priceGt, err := strconv.ParseFloat(price, 64)
		if err != nil {
			http.Error(w, "priceGt must be a valid number", http.StatusBadRequest)
			return
		}

		var filteredProducts []Product
		for _, product := range products {
			if product.Price > priceGt {
				filteredProducts = append(filteredProducts, product)
			}
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type", "appliation/json")
		json.NewEncoder(w).Encode(filteredProducts)
	})

	http.ListenAndServe(":8080", router)
}
