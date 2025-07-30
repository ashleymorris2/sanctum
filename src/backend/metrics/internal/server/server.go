package server

import (
	"github.com/labstack/echo/v4"
	"metrics/internal/server/routes"
	"metrics/internal/validators"
)

func New() *echo.Echo {
	e := echo.New()
	e.Validator = validators.NewRequestValidator()

	api := e.Group("/api")
	routes.RegisterAuth(api)
	routes.RegisterMetrics(api)

	return e
}
