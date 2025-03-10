package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"testing"

	"github.com/Sankhay/go-api-fetcher/models"
	"github.com/Sankhay/go-api-fetcher/tests"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAll(t *testing.T) {
	t.Run("GetUserByIdControllers", TestGetUserByIdControllers)

	t.Run("CreateUserControllers", TestCreateUser)
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
	w, _ := tests.PerformTestGetRequest(r, "/users/1")

	assert.Equal(t, http.StatusOK, w.Code)
}

func testGetUserByIdControllersValidationsInvalidIdRange(t *testing.T, r *gin.Engine) {
	w, body := tests.PerformTestGetRequest(r, fmt.Sprintf("/users/%s", strconv.Itoa(minIdRange-1)))
	w2, body2 := tests.PerformTestGetRequest(r, fmt.Sprintf("/users/%s", strconv.Itoa(maxIdRange+1)))

	httpError := models.HttpError{Code: http.StatusBadRequest, Msg: fmt.Sprintf(`id must be between %s and %s`, strconv.Itoa(minIdRange), strconv.Itoa(maxIdRange))}
	httpErrorJson, _ := json.Marshal(httpError)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.JSONEq(t, string(httpErrorJson), body)

	assert.Equal(t, http.StatusBadRequest, w2.Code)
	assert.JSONEq(t, string(httpErrorJson), body2)
}

func testGetUserByIdControllersValidationsInvalidIdType(t *testing.T, r *gin.Engine) {
	w, body := tests.PerformTestGetRequest(r, "/users/test")

	assert.Equal(t, http.StatusBadRequest, w.Code)
	httpError := models.HttpError{Code: http.StatusBadRequest, Msg: fmt.Sprintf(`id must be a int from %s to %s`, strconv.Itoa(minIdRange), strconv.Itoa(maxIdRange))}
	httpErrorJson, _ := json.Marshal(httpError)
	assert.JSONEq(t, string(httpErrorJson), body)
}

func TestCreateUser(t *testing.T) {
	r := tests.SetupTestMode()
	r.POST("/users/create", CreateUserControllers)

	t.Run("OK", func(t *testing.T) {
		testCreateUserOK(t, r)
	})

	t.Run("ValidationsNameAndNickNameEquals", func(t *testing.T) {
		testCreateUserValidationsNameAndNicknameEquals(t, r)
	})

	t.Run("ValidationsEmailCorrect", func(t *testing.T) {
		testCreateUserValidationsEmail(t, r)
	})

	t.Run("ValidationsNicknameNecessary", func(t *testing.T) {
		testCreateUserValidationsNicknameNecessary(t, r)
	})

	t.Run("ValidationsNameNecessary", func(t *testing.T) {
		testCreateUserValidationsNameNecessary(t, r)
	})

	t.Run("ValidationsEmailNecessary", func(t *testing.T) {
		testCreateUserValidationsEmailNecessary(t, r)
	})
}

func testCreateUserOK(t *testing.T, r *gin.Engine) {
	userTest := CreateUser{Name: "test name", Nickname: "test nickname", Email: "testEmail@test.com"}
	w, _ := tests.PerformTestPostRequest(r, "/users/create", userTest)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func testCreateUserValidationsNameAndNicknameEquals(t *testing.T, r *gin.Engine) {
	userTest := CreateUser{Name: "testEquals", Nickname: "testEquals", Email: "testEmail@test.com"}
	w, body := tests.PerformTestPostRequest(r, "/users/create", userTest)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	httpError := models.HttpError{Code: http.StatusBadRequest, Msg: "Name and nickname cannot be the same"}
	httpErrorJson, _ := json.Marshal(httpError)
	assert.JSONEq(t, string(httpErrorJson), body)
}

func testCreateUserValidationsEmail(t *testing.T, r *gin.Engine) {
	userTest := CreateUser{Name: "test name", Nickname: "test nickname", Email: "test.com"}
	w, body := tests.PerformTestPostRequest(r, "/users/create", userTest)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	httpError := models.HttpError{Code: http.StatusBadRequest, Msg: "The field 'Email' must be a valid email address. "}
	httpErrorJson, _ := json.Marshal(httpError)
	assert.JSONEq(t, string(httpErrorJson), body)
}

func testCreateUserValidationsNicknameNecessary(t *testing.T, r *gin.Engine) {
	userTest := CreateUser{Name: "test name", Email: "test@test.com"}
	w, body := tests.PerformTestPostRequest(r, "/users/create", userTest)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	httpError := models.HttpError{Code: http.StatusBadRequest, Msg: "The field 'Nickname' is required. "}
	httpErrorJson, _ := json.Marshal(httpError)
	assert.JSONEq(t, string(httpErrorJson), body)
}

func testCreateUserValidationsNameNecessary(t *testing.T, r *gin.Engine) {
	userTest := CreateUser{Nickname: "test nickname", Email: "test@test.com"}
	w, body := tests.PerformTestPostRequest(r, "/users/create", userTest)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	httpError := models.HttpError{Code: http.StatusBadRequest, Msg: "The field 'Name' is required. "}
	httpErrorJson, _ := json.Marshal(httpError)
	assert.JSONEq(t, string(httpErrorJson), body)
}

func testCreateUserValidationsEmailNecessary(t *testing.T, r *gin.Engine) {
	userTest := CreateUser{Name: "test name", Nickname: "test nickname"}
	w, body := tests.PerformTestPostRequest(r, "/users/create", userTest)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	httpError := models.HttpError{Code: http.StatusBadRequest, Msg: "The field 'Email' is required. "}
	httpErrorJson, _ := json.Marshal(httpError)
	assert.JSONEq(t, string(httpErrorJson), body)
}
