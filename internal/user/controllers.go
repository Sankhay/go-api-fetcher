package user

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Sankhay/go-api-fetcher/models"
	"github.com/gin-gonic/gin"
)

func GetUserByIdControllers(c *gin.Context) {
	userId := c.Param("id")

	userIdInt, err := strconv.Atoi(userId)

	if err != nil {
		c.JSON(http.StatusBadRequest, &models.HttpError{Code: http.StatusBadRequest, Msg: "id must be a int from 1 to 10"})
		return
	}

	if userIdInt < minIdRange || userIdInt > maxIdRange {
		c.JSON(http.StatusBadRequest, &models.HttpError{Code: http.StatusBadRequest, Msg: "id must be between 1 and 10"})
		return
	}

	user, networkError := getUserByIdServices(userId)

	if networkError != nil {
		c.JSON(networkError.Code, networkError)
		return
	}

	c.JSON(http.StatusOK, user)
}

func CreateUserControllers(c *gin.Context) {

	var createUser CreateUser

	if err := c.ShouldBindJSON(&createUser); err != nil {

		validationErrors := generateCreateUserValidationErrors(err)

		c.JSON(http.StatusBadRequest, &models.HttpError{Code: http.StatusBadRequest, Msg: validationErrors})
		return
	}

	if createUser.Name == createUser.Nickname {
		c.JSON(http.StatusBadRequest, &models.HttpError{Code: http.StatusBadRequest, Msg: "Name and nickname cannot be the same"})
		return
	}

	userCreated, networkError := createUserServices(createUser)

	if networkError != nil {
		c.JSON(networkError.Code, networkError)
		return
	}

	var userCreatedResponse struct {
		User CreateUser `json:"user" binding:"required"`
		Link string     `json:"link" binding:"required"`
	}

	userCreatedResponse.User = *userCreated
	userCreatedResponse.Link = fmt.Sprintf(`%s/users/%s`, jsonPlaceholderApiLink, strconv.Itoa(userCreated.Id))

	c.JSON(http.StatusCreated, userCreatedResponse)
}
