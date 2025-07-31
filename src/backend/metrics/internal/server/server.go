package server

import (
	"context"
	"database/sql"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"log"
	"metrics/internal/db"
	"metrics/internal/db/sqlc"
	"metrics/internal/server/routes"
	"metrics/internal/validators"
	"time"
)

type Server struct {
	Echo *echo.Echo
	DB   *sql.DB
}

func New() *Server {
	e := echo.New()
	e.Validator = validators.NewRequestValidator()

	dbCtx, dbCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer dbCancel()

	conn, err := db.ConnectToPostgres(dbCtx, db.DefaultPostgresConfig())
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
	err := s.DB.Close()
	if err != nil {
		log.Fatal(err)
	}
}
