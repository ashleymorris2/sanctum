package routes

import (
	"metrics/internal/auth"
	"metrics/internal/handlers"

	"github.com/labstack/echo/v4"
)

func RegisterAuthFor(e *echo.Group, provider auth.CredentialService) {
	h := handlers.NewAuthHandler(provider)

	authGroup := e.Group("/auth")

	authGroup.POST("/login", h.Login)
	authGroup.POST("/verify", h.VerifyAccessToken)
	authGroup.POST("/refresh", h.RefreshAccessToken)
}
