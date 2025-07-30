package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type PocketBaseProvider struct {
	BaseURL string
}

func NewPocketBaseAuth(baseUrl string) *PocketBaseProvider {
	return &PocketBaseProvider{
		BaseURL: baseUrl,
	}
}

func (pb *PocketBaseProvider) Authenticate(email, password string) (*Result, error) {
	payload := map[string]string{
		"identity": email,
		"password": password,
	}
	body, _ := json.Marshal(payload)

	res, err := http.Post(
		fmt.Sprintf("%s/api/collections/users/auth-with-password", pb.BaseURL),
		"application/json",
		bytes.NewReader(body),
	)
	if err != nil || res.StatusCode != 200 {
		return nil, echo.NewHTTPError(http.StatusUnauthorized, "Invalid email or password")
	}
	defer res.Body.Close()
	var parsed struct {
		Token  string `json:"token"`
		Record struct {
			ID    string `json:"id"`
			Email string `json:"email"`
		} `json:"record"`
	}
	if err := json.NewDecoder(res.Body).Decode(&parsed); err != nil {
		return nil, err
	}

	return &Result{
		UserID: parsed.Record.ID,
		Email:  parsed.Record.Email,
		Token:  parsed.Token,
	}, nil
}
