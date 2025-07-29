package server

import (
	"github.com/labstack/echo/v4"
	"metrics/internal/server/routes"
)

func New() *echo.Echo {
	e := echo.New()

	api := e.Group("/api")
	routes.RegisterAuth(api)
	routes.RegisterMetrics(api)

	return e
}
