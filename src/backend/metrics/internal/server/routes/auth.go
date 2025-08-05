package routes

import (
	"github.com/labstack/echo/v4"
	"metrics/internal/auth"
	"metrics/internal/handler"
)

func RegisterAuthFor(g *echo.Group, provider auth.Provider) {
	h := handler.NewAuth(provider)

	g.POST("/login", h.Login)
	g.POST("/verify", h.VerifyAuthToken)
	g.POST("/refresh", h.RefreshAuthToken)
}
