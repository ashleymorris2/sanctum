package auth

import (
	"bytes"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
)

type PocketBaseProvider struct {
}

func NewPocketBaseAuth() *PocketBaseProvider {
	return &PocketBaseProvider{
		//BaseURL: os.Getenv("AUTH_SERVICE_URL"),
	}
}

func (pb *PocketBaseProvider) Authenticate(email, password string) error {
	var request Request
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON")
	}

	if err := c.Validate(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	body, _ := json.Marshal(map[string]string{
		"identity": request.Email,
		"password": request.Password,
	})

	res, err := http.Post("http://pocketbase:8090/api/collections/users/auth-with-password", "application/json", bytes.NewBuffer(body))
	if err != nil || res.StatusCode != 200 {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid email or password")
	}
	defer res.Body.Close()

	return nil
}
