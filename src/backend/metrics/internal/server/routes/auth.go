package routes

import (
	"github.com/labstack/echo/v4"
	"metrics/internal/auth"
	"metrics/internal/db/sqlc"
	"metrics/internal/handlers"
	"os"
)

func RegisterAuthFor(g *echo.Group, queries *sqlc.Queries) {
	handler := handlers.NewAuth(auth.ByCredentials(queries, []byte(os.Getenv("JWT_SECRET"))))
	g.POST("/login", handler.Login)
}
