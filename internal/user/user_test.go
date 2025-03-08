package user

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Sankhay/go-api-fetcher/tests"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAll(t *testing.T) {
	t.Run("GetUserByIdControllers", TestGetUserByIdControllers)
}

func TestGetUserByIdControllers(t *testing.T) {
	r := tests.SetupTestMode()
	r.GET("/users/:id", GetUserByIdControllers)

	t.Run("OK", func(t *testing.T) {
		testGetUserByIdControllersOK(t, r)
	})
	t.Run("ValidationsIdRange", func(t *testing.T) {
		testGetUserByIdControllersValidationsInvalidIdRange(t, r)
	})
	t.Run("ValidationsIdType", func(t *testing.T) {
		testGetUserByIdControllersValidationsInvalidIdType(t, r)
	})
}

func testGetUserByIdControllersOK(t *testing.T, r *gin.Engine) {

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/users/1", nil)

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func testGetUserByIdControllersValidationsInvalidIdRange(t *testing.T, r *gin.Engine) {

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/users/0", nil)

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.JSONEq(t, `{"message": "id must be between 1 and 10"}`, w.Body.String())
}

func testGetUserByIdControllersValidationsInvalidIdType(t *testing.T, r *gin.Engine) {

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/users/test", nil)

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.JSONEq(t, `{"message": "id must be a int from 1 to 10"}`, w.Body.String())
}
