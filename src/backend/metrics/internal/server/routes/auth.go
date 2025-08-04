package routes

import (
	"github.com/labstack/echo/v4"
	"metrics/internal/auth"
	"metrics/internal/handlers"
)

func RegisterAuthFor(g *echo.Group, provider auth.Provider) {
	handler := handlers.NewAuth(provider)
	g.POST("/login", handler.Login)
}
