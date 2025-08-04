package server

import (
	"context"
	"database/sql"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"log"
	"metrics/internal/auth"
	"metrics/internal/db"
	"metrics/internal/db/repositories"
	"metrics/internal/db/sqlc"
	"metrics/internal/middleware"
	"metrics/internal/server/routes"
	"metrics/internal/validators"
	"os"
	"time"
)

type Server struct {
	Echo *echo.Echo
	DB   *sql.DB
}

func New() *Server {
	conn := dbConect()
	queries := sqlc.New(conn)

	authProvider := configureAuth(queries)

	e := echo.New()
	e.Validator = validators.NewRequestValidator()

	e.Use(middleware.AuthMiddleware(auth.NewMiddlewareConfig(authProvider)))

	api := e.Group("/api")
	routes.RegisterAuthFor(api, authProvider)
	routes.RegisterMetricsFor(api)

	return &Server{
		Echo: e,
		DB:   conn,
	}
}

func configureAuth(queries *sqlc.Queries) auth.Provider {
	return auth.ByCredentials(
		queries,
		*repositories.NewRefreshTokenRepository(queries),
		[]byte(os.Getenv("JWT_SECRET")))
}

func dbConect() *sql.DB {
	dbCtx, dbCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer dbCancel()

	conn, err := db.ConnectToPostgres(dbCtx, db.DefaultPostgresConfig())
	if err != nil {
		log.Fatal(err)
	}
	return conn
}

func (s *Server) Shutdown() {
	err := s.DB.Close()
	if err != nil {
		log.Fatal(err)
	}
}
