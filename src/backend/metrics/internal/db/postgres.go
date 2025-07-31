package db

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"
)

// PostgresConfig holds database connection parameters
type PostgresConfig struct {
	// URL is the database connection string. If empty, will use AUTH_DATABASE_URL env var
	URL string
	// MaxOpenConns controls the maximum number of open connections to the database
	MaxOpenConns int
	// MaxIdleConns controls the maximum number of idle connections in the pool
	MaxIdleConns int
	// ConnMaxLifetime controls the maximum amount of time a connection may be reused
	ConnMaxLifetime time.Duration
}

// DefaultPostgresConfig returns a PostgresConfig with reasonable defaults
func DefaultPostgresConfig() PostgresConfig {
	return PostgresConfig{
		MaxOpenConns:    25,
		MaxIdleConns:    5,
		ConnMaxLifetime: 5 * time.Minute,
	}
}

// ConnectToPostgres establishes a connection to a PostgreSQL database with the provided config
// Returns a database connection and any error encountered
func ConnectToPostgres(ctx context.Context, config PostgresConfig) (*sql.DB, error) {
	dbURL := config.URL
	if dbURL == "" {
		dbURL = os.Getenv("AUTH_DATABASE_URL")
		if dbURL == "" {
			return nil, fmt.Errorf("database URL not provided and AUTH_DATABASE_URL env var is empty")
		}
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	// Configure the connection pool
	conn.SetMaxOpenConns(config.MaxOpenConns)
	conn.SetMaxIdleConns(config.MaxIdleConns)
	conn.SetConnMaxLifetime(config.ConnMaxLifetime)

	// Verify the connection works
	err = conn.PingContext(ctx)
	if err != nil {
		conn.Close() // Close the connection before returning the error
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return conn, nil
}
