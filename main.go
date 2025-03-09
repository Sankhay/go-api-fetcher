package main

import (
	"fmt"

	"github.com/Sankhay/go-api-fetcher/config"
	"github.com/Sankhay/go-api-fetcher/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	config.Init()

	r := gin.Default()

	routes.StartRoutes(r)

	r.Run(fmt.Sprintf(":%s", config.Port))
}
