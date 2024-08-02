package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/pietroBragaAquinoJunior/go-gin-gorm-crud/src/models"
	"github.com/pietroBragaAquinoJunior/go-gin-gorm-crud/src/router"
	"github.com/stretchr/testify/assert"
)

func TestGetProductById(t *testing.T) {
	router := router.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/product/10", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "10", w.Body.String())
}

func TestPostUser(t *testing.T) {
	router := router.SetupRouter()

	exampleProduct := models.Product{
		Name:  "Desodorante",
		Price: 19.90,
	}

	productJson, _ := json.Marshal(exampleProduct)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/product/add", strings.NewReader(string(productJson)))
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, string(productJson), w.Body.String())
}
