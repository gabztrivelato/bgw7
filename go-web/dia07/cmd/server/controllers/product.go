package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gabztrivelato/bgw7/dia07/internal/domain"
	"github.com/gabztrivelato/bgw7/dia07/internal/product"
	"github.com/go-chi/chi/v5"
)

type ProductController struct {
	repository product.Repository
}

func NewProductController(repository product.Repository) *ProductController {
	return &ProductController{repository: repository}
}

func (c *ProductController) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var productReq domain.ProductRequest
	err := json.NewDecoder(r.Body).Decode(&productReq)
	if err != nil {
		http.Error(w, "error processing json body", http.StatusUnprocessableEntity)
		return
	}

	err = productReq.Validate()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := c.repository.Create(productReq.ToDomain())
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	response := domain.CreateProductResponse{Id: id}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (c *ProductController) GetById(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	idParam := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(idParam)
	product := c.repository.GetById(id)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

func (c *ProductController) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var productReq domain.ProductRequest
	err = json.NewDecoder(r.Body).Decode(&productReq)
	if err != nil {
		http.Error(w, "error processing json body", http.StatusUnprocessableEntity)
		return
	}

	err = productReq.Validate()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	product := productReq.ToDomain()
	product.Id = id

	err = c.repository.Update(product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)

}

func (c *ProductController) Patch(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var productReq domain.PatchProductRequest
	err = json.NewDecoder(r.Body).Decode(&productReq)
	if err != nil {
		http.Error(w, "error processing json body", http.StatusUnprocessableEntity)
		return
	}

	err = c.repository.PartialUpdate(id, productReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(productReq)

}

func (c *ProductController) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.repository.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
