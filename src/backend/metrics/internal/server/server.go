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
	conn := dbConnect()
	queries := sqlc.New(conn)
	authService := configureAuth(queries)

	e := echo.New()
	e.Validator = validators.NewRequestValidator()

	public := e.Group("/api")
	routes.RegisterAuthFor(public, authService)

	private := e.Group("/api")
	private.Use(middleware.AuthMiddleware(auth.NewMiddlewareConfig(authService)))
	routes.RegisterMetricsFor(private)

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

func dbConnect() *sql.DB {
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
