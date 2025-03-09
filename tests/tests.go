package tests

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func SetupTestMode() *gin.Engine {
	r := gin.Default()

	_, filename, _, _ := runtime.Caller(0)

	absPath := filepath.Dir(filename)

	err := godotenv.Load(absPath + "/tests.env")

	if err != nil {
		log.Fatal("Error loading tests.env file, it should be in tests/, ", err)
	}

	gin.SetMode(gin.TestMode)

	return r
}

func PerformTestPostRequest(r *gin.Engine, endpoint string, payload interface{}) (*httptest.ResponseRecorder, string) {
	jsonPayload, _ := json.Marshal(payload)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonPayload))
	r.ServeHTTP(w, req)
	return w, w.Body.String()
}

func PerformTestGetRequest(r *gin.Engine, endpoint string) (*httptest.ResponseRecorder, string) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", endpoint, nil)
	r.ServeHTTP(w, req)
	return w, w.Body.String()
}
