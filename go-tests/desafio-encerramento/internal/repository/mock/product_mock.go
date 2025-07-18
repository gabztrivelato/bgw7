package repository

import (
	"app/internal"

	"github.com/stretchr/testify/mock"
)

type ProductRepositoryMock struct {
	mock.Mock
}

func (r *ProductRepositoryMock) SearchProducts(query internal.ProductQuery) (map[int]internal.Product, error) {
	args := r.Mock.Called()
	return args.Get(0).(map[int]internal.Product), args.Error(1)
}
