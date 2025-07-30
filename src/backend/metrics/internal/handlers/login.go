package handlers

import (
	"github.com/labstack/echo/v4"
	"metrics/internal/auth"
	"net/http"
)

var authService auth.Provider = auth.NewPocketBaseAuth()

func Login(c echo.Context) error {
	var request auth.Request
	if err := c.Bind(&request); err != nil || c.Validate(&request) != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid input")
	}

	authResult, err := authService.Authenticate(request.Email, request.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid credentials")
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token":  authResult.Token,
		"userId": authResult.UserID,
	})
}
