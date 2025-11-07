package routes

import (
	"metrics/internal/auth"
	"metrics/internal/handlers"

	"github.com/labstack/echo/v4"
)

func RegisterAuthFor(g *echo.Group, provider auth.CredentialService) {
	h := handlers.NewAuthHandler(provider)

	authGroup := g.Group("/auth")

	authGroup.POST("/login", h.Login)
	authGroup.POST("/verify", h.VerifyAuthToken)
	authGroup.POST("/refresh", h.RefreshAccessToken)
}
