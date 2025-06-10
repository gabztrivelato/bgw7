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
	Create(product domain.Product) (int, error)
	GetById(id int) domain.Product
}

func NewRepository() Repository {
	return &repository{}
}

func LoadData() {
	f, err := os.Open("../../docs/products.json")
	if err != nil {
		panic(err)
	}

	defer f.Close()
	var products map[int]domain.Product
	err = json.NewDecoder(f).Decode(&products)
	if err != nil {
		panic(err)
	}
}

func (r *repository) Create(product domain.Product) (int, error) {

	for _, productStoraged := range r.storage {
		if productStoraged.Code == product.Code {
			return 0, errors.New("product code already exists")
		}
	}
	r.lastId++

	r.storage[r.lastId] = product
	return r.lastId, nil
}

func (r *repository) GetById(id int) domain.Product {
	product := r.storage[id]
	return product
}
