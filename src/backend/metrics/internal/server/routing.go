package server

import (
	"metrics/internal/auth"
	"metrics/internal/routes"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func registerWellKnown(e *echo.Echo) {
	routes.RegisterWellKnownRoutes(e)
}

func registerSwagger(e *echo.Echo) {
	if showSwagger {
		e.GET("/swagger/*", echoSwagger.WrapHandler)
	}
}

func registerPublicRoutes(e *echo.Group, authService auth.CredentialService) {
	routes.RegisterAuthFor(e, authService)
}

func registerRoutes(e *echo.Group, authService auth.CredentialService) {
	routes.RegisterMetricsFor(e, authService)
}
