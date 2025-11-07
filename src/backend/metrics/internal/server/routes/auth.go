package routes

import (
	"metrics/internal/auth"
	"metrics/internal/handler"

	"github.com/labstack/echo/v4"
)

func RegisterAuthFor(g *echo.Group, provider auth.CredentialService) {
	h := handler.NewAuthHandler(provider)

	authGroup := g.Group("/auth")

	authGroup.POST("/login", h.Login)
	authGroup.POST("/verify", h.VerifyAuthToken)
	authGroup.POST("/refresh", h.RefreshAuthToken)
}
