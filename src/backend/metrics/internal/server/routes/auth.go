package routes

import (
	"github.com/labstack/echo/v4"
	"metrics/internal/handlers"
)

func RegisterAuth(g *echo.Group) {
	g.POST("/login", handlers.Login)
}
