package main

import (
	"log"

	"github.com/Sankhay/go-api-fetcher/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	r := gin.Default()

	routes.StartRoutes(r)

	r.Run(":8000")
}
