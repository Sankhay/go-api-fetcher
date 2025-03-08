package user

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/Sankhay/go-api-fetcher/models"
)

func getUserByIdService(userId string) (*User, *models.NetworkError) {
	link := fmt.Sprintf("%s/users/%s", jsonPlaceholderApiLink, userId)

	resp, err := http.Get(link)

	if err != nil {
		log.Println(err)
		return nil, &models.NetworkError{Code: http.StatusInternalServerError, Msg: err.Error()}
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Println(err)
		return nil, &models.NetworkError{Code: http.StatusInternalServerError, Msg: err.Error()}
	}

	var user User

	if err := json.Unmarshal(body, &user); err != nil {
		log.Println(err)
		return nil, &models.NetworkError{Code: http.StatusInternalServerError, Msg: err.Error()}
	}

	return &user, nil
}
