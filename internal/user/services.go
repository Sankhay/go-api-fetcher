package user

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/Sankhay/go-api-fetcher/models"
)

func getUserByIdServices(userId string) (*User, *models.HttpError) {
	link := fmt.Sprintf("%s/users/%s", jsonPlaceholderApiLink, userId)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, link, nil)

	if err != nil {
		errMsg := fmt.Sprintf("Failed to create request for user id %s: %v", userId, err)
		log.Println(errMsg)
		return nil, &models.HttpError{Code: http.StatusInternalServerError, Msg: errMsg}
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		errMsg := fmt.Sprintf("Request failed for user id %s: %v", userId, err)
		log.Println(errMsg)
		return nil, &models.HttpError{Code: http.StatusInternalServerError, Msg: errMsg}
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		errMsg := fmt.Sprintf("Request failed for user id %s: %v", userId, err)
		log.Println(errMsg)
		return nil, &models.HttpError{Code: http.StatusInternalServerError, Msg: errMsg}
	}

	if resp.StatusCode != http.StatusOK {
		// I dont found any error returning with a body
		errMsg := fmt.Sprintf("Unexpected status code %d for user id %s: %s", resp.StatusCode, userId, string(body))
		log.Println(errMsg)
		return nil, &models.HttpError{Code: resp.StatusCode, Msg: errMsg}
	}

	var user User

	if err := json.Unmarshal(body, &user); err != nil {
		errMsg := fmt.Sprintf("Failed to parse json response for user id %s: %v", userId, err)
		log.Println(errMsg)
		return nil, &models.HttpError{Code: http.StatusInternalServerError, Msg: errMsg}
	}

	return &user, nil
}

func createUserServices(createUser CreateUser) (*CreateUser, *models.HttpError) {
	link := fmt.Sprintf("%s/users", jsonPlaceholderApiLink)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	jsonCreateUserData, err := json.Marshal(createUser)
	if err != nil {
		errMsg := fmt.Sprintf("Failed to marshal CreateUser data: %v", err)
		log.Println(errMsg)
		return nil, &models.HttpError{Code: http.StatusInternalServerError, Msg: errMsg}
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, link, bytes.NewBuffer(jsonCreateUserData))

	if err != nil {
		errMsg := fmt.Sprintf("Failed to create request for creating user: %v", err)
		log.Println(errMsg)
		return nil, &models.HttpError{Code: http.StatusInternalServerError, Msg: errMsg}
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		errMsg := fmt.Sprintf("Request failed to create user: %v", err)
		log.Println(errMsg)
		return nil, &models.HttpError{Code: http.StatusInternalServerError, Msg: errMsg}
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		errMsg := fmt.Sprintf("Failed to read response body when creating user: %v", err)
		log.Println(errMsg)
		return nil, &models.HttpError{Code: http.StatusInternalServerError, Msg: errMsg}
	}

	if resp.StatusCode != http.StatusCreated {
		// I dont found any error returning with a body
		errMsg := fmt.Sprintf("API returned unexpected status code %d while creating user , error api body: %s", resp.StatusCode, string(body))
		log.Println(errMsg)
		return nil, &models.HttpError{Code: resp.StatusCode, Msg: errMsg}
	}

	var createdUser CreateUser

	if err := json.Unmarshal(body, &createdUser); err != nil {
		errMsg := fmt.Sprintf("Failed to parse json response while creating user: %v", err)
		log.Println(errMsg)
		return nil, &models.HttpError{Code: http.StatusInternalServerError, Msg: errMsg}
	}

	return &createdUser, nil
}
