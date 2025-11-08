package routes

import (
	"log"
	"metrics/internal/auth"
	"os"

	"github.com/labstack/echo/v4"
)

func RegisterWellKnownRoutes(e *echo.Echo) {
	g := e.Group("/.well-known")

	keys, err := auth.LoadKeyPairFromPEM(os.Getenv("JWT_RSA_PRIV_PATH"), os.Getenv("JWT_KID"))
	if err != nil {
		log.Fatal(err)
	}

	g.GET("/jwks.json", auth.JWKSHandler())
}
