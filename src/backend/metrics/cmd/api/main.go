package main

import (
	"github.com/joho/godotenv"
	"log"
	"metrics/internal/server"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Warning: Error loading .env file: %v", err)
	}

	s := server.New()
	defer s.Shutdown()

	log.Fatal(s.Echo.Start("0.0.0.0:4200"))
}
