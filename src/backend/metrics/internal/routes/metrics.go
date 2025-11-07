package routes

import "github.com/labstack/echo/v4"

func RegisterMetricsFor(g *echo.Group) {
	metricsGroup := g.Group("/metricsGroup")

	metricsGroup.GET("/", nil)
}
