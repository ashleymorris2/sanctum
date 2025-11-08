package routes

import (
	"metrics/internal/auth"
	"metrics/internal/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterMetricsFor(g *echo.Group, authService auth.CredentialService) {
	metricsGroup := g.Group("/metrics")
	metricsGroup.Use(middleware.AuthMiddleware(auth.NewMiddlewareConfig(authService)))

	metricsGroup.GET("/", nil)
}
