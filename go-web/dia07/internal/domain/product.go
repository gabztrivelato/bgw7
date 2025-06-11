package domain

import (
	"errors"
	"regexp"
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

type ProductRequest struct {
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	Code        string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

type PatchProductRequest struct {
	Name        string   `json:"name"`
	Quantity    *int     `json:"quantity"`
	Code        string   `json:"code_value"`
	IsPublished *bool    `json:"is_published"`
	Expiration  string   `json:"expiration"`
	Price       *float64 `json:"price"`
}

type CreateProductResponse struct {
	Id int `json:"id"`
}

func (product *ProductRequest) ToDomain() Product {
	return Product{
		Name:        product.Name,
		Quantity:    product.Quantity,
		Code:        product.Code,
		IsPublished: product.IsPublished,
		Expiration:  product.Expiration,
		Price:       product.Price,
	}
}

func (product *ProductRequest) Validate() error {
	if product.Name == "" {
		return errors.New("name must have a value")
	}

	if product.Quantity == 0 {
		return errors.New("quantity must have a value")
	}

	if product.Code == "" {
		return errors.New("code must have a value")
	}

	dateRegex := regexp.MustCompile(`^\d{2}\/(0[1-9]|1[0-2])\/\d{4}$`)
	if !dateRegex.MatchString(product.Expiration) {
		return errors.New("expiration date must be dd/MM/YYYY")
	}

	if product.Price == 0.0 {
		return errors.New("price must have a value")
	}
	return nil
}
