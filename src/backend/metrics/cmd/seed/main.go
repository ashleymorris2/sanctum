package main

import (
	"context"
	"database/sql"

	"fmt"
	_ "github.com/lib/pq"
	"log"
	"metrics/internal/auth"
	"metrics/internal/db/sqlc"

	"os"
)

func main() {
	connStr := os.Getenv("AUTH_DATABASE_URL")
	if connStr == "" {
		log.Fatal("AUTH_DATABASE_URL not set")
	}

	dbConn, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer dbConn.Close()
	queries := sqlc.New(dbConn)

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET not set")
	}

	credentials := auth.ByCredentials(queries, []byte(jwtSecret))

	ctx := context.Background()
	// Check if user exists
	_, err = credentials.Authenticate(ctx, "user@test.com", "password")
	if err == nil {
		fmt.Println("Seed user already exists.")
		return
	}

	// Create user
	res, err := credentials.Register(ctx, "user@test.com", "password")
	if err != nil {
		log.Fatalf("Failed to seed user: %v", err)
	}

	fmt.Printf(" User %s seeded with ID: %s\n", res.Email, res.UserID)
}
