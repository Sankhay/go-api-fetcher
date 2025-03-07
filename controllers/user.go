package controllers

import (
	"log"
	"net/http"

	"github.com/Sankhay/go-api-fetcher/api"
	"github.com/gin-gonic/gin"
)

func GetUserById(c *gin.Context) {
	userId := c.Param("id")

	user, err := api.GetUserById(userId)

	if err != nil {
		log.Fatalln(err)
	}

	c.JSON(http.StatusOK, user)
}
