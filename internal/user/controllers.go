package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUserByIdControllers(c *gin.Context) {
	userId := c.Param("id")

	userIdInt, err := strconv.Atoi(userId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "id must be a int from 1 to 10",
		})
		return
	}

	if userIdInt < minIdRange || userIdInt > maxIdRange {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "id must be between 1 and 10",
		})
		return
	}

	user, networkError := getUserByIdService(userId)

	if networkError != nil {
		c.JSON(networkError.Code, gin.H{
			"message": networkError.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, user)
}
