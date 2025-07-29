package routes

import "github.com/labstack/echo/v4"

func RegisterMetrics(g *echo.Group) {
	metrics := g.Group("/metrics")
	metrics.GET("/", nil)
}
