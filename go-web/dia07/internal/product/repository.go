package product

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/gabztrivelato/bgw7/dia07/internal/domain"
)

type repository struct {
	storage map[int]domain.Product
	lastId  int
}

type Repository interface {
	LoadData()
	Create(product domain.Product) (int, error)
	GetById(id int) *domain.Product
	Update(product domain.Product) error
	Delete(id int) error
	PartialUpdate(id int, product domain.PatchProductRequest) error
}

func NewRepository() Repository {
	return &repository{
		storage: make(map[int]domain.Product),
	}
}

func (r *repository) LoadData() {
	f, err := os.Open("../../docs/products.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var products []domain.Product
	err = json.NewDecoder(f).Decode(&products)
	if err != nil {
		panic(err)
	}
	for _, product := range products {
		if product.Id > r.lastId {
			r.lastId = product.Id
		}
		r.storage[product.Id] = product
	}
}

func (r *repository) Create(product domain.Product) (int, error) {
	for _, productStoraged := range r.storage {
		if productStoraged.Code == product.Code {
			return 0, errors.New("product code already exists")
		}
	}
	r.lastId++

	product.Id = r.lastId
	r.storage[r.lastId] = product

	return r.lastId, nil
}

func (r *repository) GetById(id int) *domain.Product {
	product, ok := r.storage[id]
	if !ok {
		return nil
	}
	return &product
}

func (r *repository) Update(product domain.Product) error {
	_, ok := r.storage[product.Id]
	if !ok {
		return errors.New("product not found")
	}
	r.storage[product.Id] = product
	return nil
}

func (r *repository) PartialUpdate(id int, product domain.PatchProductRequest) error {
	productStoraged := r.GetById(id)

	if productStoraged == nil {
		return errors.New("product not found")
	}

	if product.Name != "" {
		productStoraged.Name = product.Name
	}

	if product.Quantity != nil {
		productStoraged.Quantity = *product.Quantity
	}

	if product.Code != "" {
		productStoraged.Code = product.Code
	}

	if product.IsPublished != nil {
		productStoraged.IsPublished = *product.IsPublished
	}

	if product.Expiration != "" {
		productStoraged.Expiration = product.Expiration
	}
	if product.Price != nil {
		productStoraged.Price = *product.Price
	}

	r.storage[id] = *productStoraged
	return nil
}

func (r *repository) Delete(id int) error {
	_, ok := r.storage[id]
	if !ok {
		return errors.New("product not found")
	}
	delete(r.storage, id)
	return nil
}
