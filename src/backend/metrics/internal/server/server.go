package server

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"log"
	"metrics/internal/db/sqlc"
	"metrics/internal/server/routes"
	"metrics/internal/validators"
	"os"
)

type Server struct {
	Echo *echo.Echo
	DB   *sql.DB
}

func New() *Server {
	e := echo.New()
	e.Validator = validators.NewRequestValidator()

	dbURL := os.Getenv("AUTH_DATABASE_URL")
	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}

	queries := sqlc.New(conn)

	api := e.Group("/api")
	routes.RegisterAuthFor(api, queries)
	routes.RegisterMetricsFor(api)

	return &Server{
		Echo: e,
		DB:   conn,
	}
}

func (s *Server) Shutdown() {
	s.DB.Close()
}
