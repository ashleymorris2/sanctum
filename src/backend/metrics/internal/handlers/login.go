package handlers

import (
	"errors"
	"github.com/labstack/echo/v4"
	"metrics/internal/auth"
	"net/http"
)

type RequestDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=128"`
}

type Auth struct {
	authProvider auth.Provider
}

func NewAuth(authProvider auth.Provider) *Auth {
	return &Auth{authProvider: authProvider}
}

func (a *Auth) Login(c echo.Context) error {
	ctx := c.Request().Context()

	var req RequestDTO
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, ErrorResponse{Message: "Invalid request"})
	}
	if err := c.Validate(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, ErrorResponse{
			Message: "Validation failed",
			// optionally include: "Details": err.Error(),
		})
	}

	authResult, err := a.authProvider.Authenticate(ctx, req.Email, req.Password)
	if err != nil {
		if errors.Is(err, auth.ErrAuthFailure) {
			return echo.NewHTTPError(http.StatusUnauthorized, ErrorResponse{
				Message: "Invalid credentials",
			})
		} else {
			return echo.NewHTTPError(http.StatusInternalServerError, ErrorResponse{
				Message: "Internal server error",
			})
		}
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token":  authResult.Token,
		"userId": authResult.UserID,
	})
}
