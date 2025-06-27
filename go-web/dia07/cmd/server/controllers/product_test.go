package controllers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gabztrivelato/bgw7/dia07/cmd/server/controllers"
	"github.com/gabztrivelato/bgw7/dia07/internal/domain"
	"github.com/gabztrivelato/bgw7/dia07/internal/product"
	"github.com/stretchr/testify/assert"
)

func TestGetProduct(t *testing.T) {
	t.Run("sucess to get products", func(t *testing.T) {
		db := map[int]domain.Product{
			1: {Name: "Oil - Margarine", Quantity: 439, Code: "S82254D", IsPublished: true, Expiration: "15/12/2021", Price: 71.42},
		}
		repo := product.NewRepository(db)
		ct := controllers.NewProductController(repo)

		req := httptest.NewRequest("GET", "/products/", nil)
		res := httptest.NewRecorder()
		ct.GetById(res, req)

		expectedCode := http.StatusOK
		expectedBody := `{"id": 1,"name": "Oil - Margarine","quantity": 439,"code_value":"S82254D","is_published": true,"expiration": "15/12/2021","price": 71.42}`
		assert.Equal(t, expectedCode, res.Code)
		assert.Equal(t, expectedBody, res.Body.String())
	})
}
