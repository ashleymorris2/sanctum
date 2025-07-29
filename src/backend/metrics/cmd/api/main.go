package main

import (
	"log"
	"metrics/internal/server"
)

func main() {
	app := server.New()
	log.Fatal(app.Start("0.0.0.0:4200"))
}
