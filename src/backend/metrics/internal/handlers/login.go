package handlers

import (
	"github.com/labstack/echo/v4"
	"metrics/internal/auth"
	"net/http"
)

type Auth struct {
	authProvider auth.Provider
}

func NewAuth(authProvider auth.Provider) *Auth {
	return &Auth{authProvider: authProvider}
}

func (a *Auth) Login(c echo.Context) error {
	ctx := c.Request().Context()

	var request auth.Request
	if err := c.Bind(&request); err != nil || c.Validate(&request) != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid input")
	}

	authResult, err := a.authProvider.Authenticate(ctx, request.Email, request.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid credentials")
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token":  authResult.Token,
		"userId": authResult.UserID,
	})
}
