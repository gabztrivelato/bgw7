package repository_test

import (
	"app/internal"
	"app/internal/repository"
	"testing"
)

func TestSearchProducts(t *testing.T) {
	productsAttributes := map[int]internal.ProductAttributes{
		1: {Description: "Utensílios", Price: 15.0, SellerId: 1},
		2: {Description: "Eletronicos", Price: 400.0, SellerId: 2},
		3: {Description: "Roupas", Price: 90.0, SellerId: 1},
	}

	products := map[int]internal.Product{
		1: {Id: 1, ProductAttributes: productsAttributes[1]},
		2: {Id: 2, ProductAttributes: productsAttributes[2]},
		3: {Id: 3, ProductAttributes: productsAttributes[3]},
	}
	repo := repository.NewProductsMap(products)

	query := internal.ProductQuery{Id: 2}
	result, err := repo.SearchProducts(query)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(result) != 1 || result[2].Description != "Eletronicos" {
		t.Errorf("expected to find Product 2, got %v", result)
	}
}
