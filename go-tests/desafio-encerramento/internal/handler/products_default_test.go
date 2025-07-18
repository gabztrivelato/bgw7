package handler_test

import (
	"app/internal"
	"app/internal/handler"
	repository "app/internal/repository"
	mock "app/internal/repository/mock"

	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetProduct(t *testing.T) {
	url := "/products"
	t.Run("success to get product returns 200", func(t *testing.T) {
		productsAttributes := map[int]internal.ProductAttributes{
			1: {Description: "Utensílios", Price: 15.0, SellerId: 1},
			2: {Description: "Eletronicos", Price: 400.0, SellerId: 2},
			3: {Description: "Roupas", Price: 90.0, SellerId: 1},
		}

		db := map[int]internal.Product{
			1: {Id: 1, ProductAttributes: productsAttributes[1]},
			2: {Id: 2, ProductAttributes: productsAttributes[2]},
			3: {Id: 3, ProductAttributes: productsAttributes[3]},
		}
		repository := repository.NewProductsMap(db)
		handler := handler.NewProductsDefault(repository)

		req := httptest.NewRequest("GET", url, nil)
		res := httptest.NewRecorder()

		expectedStatusCode := http.StatusOK
		expectedBody := `{"data":{"1":{"description":"Utensílios", "id":1, "price":15, "seller_id":1}, "2":{"description":"Eletronicos", "id":2, "price":400, "seller_id":2}, "3":{"description":"Roupas", "id":3, "price":90, "seller_id":1}}, "message":"success"}`
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		handler.Get()(res, req)

		assert.Equal(t, expectedStatusCode, res.Code)
		assert.JSONEq(t, expectedBody, res.Body.String())
		assert.Equal(t, expectedHeader, res.Header())
	})

	t.Run("success to get product with query id returns 200", func(t *testing.T) {
		productsAttributes := map[int]internal.ProductAttributes{
			1: {Description: "Utensílios", Price: 15.0, SellerId: 1},
			2: {Description: "Eletronicos", Price: 400.0, SellerId: 2},
			3: {Description: "Roupas", Price: 90.0, SellerId: 1},
		}
		db := map[int]internal.Product{
			1: {Id: 1, ProductAttributes: productsAttributes[1]},
			2: {Id: 2, ProductAttributes: productsAttributes[2]},
			3: {Id: 3, ProductAttributes: productsAttributes[3]},
		}
		repository := repository.NewProductsMap(db)
		handler := handler.NewProductsDefault(repository)

		req := httptest.NewRequest("GET", url+"?id=2", nil)
		res := httptest.NewRecorder()

		expectedStatusCode := http.StatusOK
		expectedBody := `{"data":{"2":{"description":"Eletronicos", "id":2, "price":400, "seller_id":2}}, "message":"success"`
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		handler.Get()(res, req)

		assert.Equal(t, expectedStatusCode, res.Code)
		assert.JSONEq(t, expectedBody, res.Body.String())
		assert.Equal(t, expectedHeader, res.Header())
	})

	t.Run("failed to get products", func(t *testing.T) {
		repository := mock.ProductRepositoryMock{}
		handler := handler.NewProductsDefault(&repository)

		req := httptest.NewRequest(http.MethodGet, url, nil)
		res := httptest.NewRecorder()

		expectedStatusCode := http.StatusInternalServerError
		expectedBody := `{"status":"Internal Server Error","message":"internal error"}`
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		var mockedProducts map[int]internal.Product
		mockedError := errors.New("failed to connect database")

		repository.
			On("SearchProducts").
			Return(mockedProducts, mockedError)

		handler.Get()(res, req)

		assert.Equal(t, expectedStatusCode, res.Code)
		assert.Equal(t, expectedBody, res.Body.String())
		assert.Equal(t, expectedHeader, res.Header())

	})
}
