package routes

import (
	"fmt"
	"github.com/labstack/echo/v4"
)

func RegisterAuth(g *echo.Group) {
	g.POST("/sign-in", func(c echo.Context) error {
		fmt.Print("sign-in called")

		return nil
	})
}
