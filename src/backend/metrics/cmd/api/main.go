package main

import (
	"log"
	"metrics/internal/server"

	_ "metrics/docs"

	"github.com/joho/godotenv"
	_ "github.com/swaggo/echo-swagger"
)

// @title			Metrics API
// @version			1.0
// @description		API for uploading of metric data
// @description		Bearer JWT (RS256). Public JWKS at /.well-known/jwks.json
// @BasePath		/api
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Warning: Error loading .env file: %v", err)
	}

	s := server.New()
	defer s.Shutdown()

	log.Fatal(s.Echo.Start("0.0.0.0:3000"))
}
