package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/pietroBragaAquinoJunior/go-gin-gorm-crud/src/controllers"
	"github.com/pietroBragaAquinoJunior/go-gin-gorm-crud/src/models"
	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	router := controllers.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}

// Test for POST /user/add
func TestPostUser(t *testing.T) {
	router := controllers.SetupRouter()
	router = controllers.PostUser(router)

	w := httptest.NewRecorder()

	// Create an example user for testing
	exampleUser := models.User{
		Username: "test_name",
		Gender:   "male",
	}
	userJson, _ := json.Marshal(exampleUser)
	req, _ := http.NewRequest("POST", "/user/add", strings.NewReader(string(userJson)))
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	// Compare the response body with the json data of exampleUser
	assert.Equal(t, string(userJson), w.Body.String())
}