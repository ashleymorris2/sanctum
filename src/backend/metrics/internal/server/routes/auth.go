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

	g.POST("/verify", h.VerifyAuthToken)
	g.POST("/refresh", h.RefreshAuthToken)
}
