package tests

import (
	"log"
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
