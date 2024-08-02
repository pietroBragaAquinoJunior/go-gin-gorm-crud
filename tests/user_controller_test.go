package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/pietroBragaAquinoJunior/go-gin-gorm-crud/src/database"
	"github.com/pietroBragaAquinoJunior/go-gin-gorm-crud/src/models"
	"github.com/pietroBragaAquinoJunior/go-gin-gorm-crud/src/router"
	"github.com/stretchr/testify/assert"
)

func TestGetProductById(t *testing.T) {
	database.Connect()
	r := router.SetupRouter()

	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/product/4", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var responseProduct models.Product
	if err := json.Unmarshal(w.Body.Bytes(), &responseProduct); err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}

	assert.Equal(t, "Desodorante", responseProduct.Name)
	assert.Equal(t, 19.90, responseProduct.Price)
}

func TestPostUser(t *testing.T) {
	database.Connect()

	r := router.SetupRouter()

	exampleProduct := models.Product{
		Name:  "Desodorante",
		Price: 19.90,
	}

	productJson, err := json.Marshal(exampleProduct)
	if err != nil {
		t.Fatalf("Failed to marshal product: %v", err)
	}

	req, err := http.NewRequest("POST", "/product/add", bytes.NewBuffer(productJson))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var actualProduct models.Product
	if err := json.Unmarshal(w.Body.Bytes(), &actualProduct); err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}

	assert.Equal(t, exampleProduct.Name, actualProduct.Name)
	assert.Equal(t, exampleProduct.Price, actualProduct.Price)
}
