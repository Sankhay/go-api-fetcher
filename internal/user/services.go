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
	link := fmt.Sprintf("https://jsonplaceholder.typicode.com/users/%s", userId)

	resp, err := http.Get(link)

	if err != nil {
		log.Println(err)
		networkError := models.NetworkError{Code: http.StatusInternalServerError, Msg: err.Error()}
		return nil, &networkError
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Println(err)
		networkError := models.NetworkError{Code: http.StatusInternalServerError, Msg: err.Error()}
		return nil, &networkError
	}

	var user User

	print(string(body))

	if err := json.Unmarshal(body, &user); err != nil {
		log.Fatal(err)
		networkError := models.NetworkError{Code: http.StatusInternalServerError, Msg: err.Error()}
		return nil, &networkError
	}

	return &user, nil
}
