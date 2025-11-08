package server

import (
	"context"
	"crypto/rsa"
	"database/sql"
	"log"
	"metrics/internal/auth"
	"metrics/internal/db"
	"metrics/internal/db/repositories"
	"metrics/internal/db/sqlc"
	"metrics/internal/middleware"
	"metrics/internal/validators"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

// showSwagger is a flag to show swagger UI
// todo: make it configurable via env
const showSwagger = true

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
	e.HTTPErrorHandler = middleware.JSONErrorHandler(e.DefaultHTTPErrorHandler)

	api := e.Group("/api")

	registerSwagger(e)
	registerWellKnown(e)

	registerPublicRoutes(api, authService)
	registerRoutes(api, authService)

	return &Server{
		Echo: e,
		DB:   conn,
	}
}

func configureAuth(queries *sqlc.Queries) auth.CredentialService {
	path := os.Getenv("JWT_RSA_PRIV_PATH")
	kid := os.Getenv("JWT_KID")

	// Load the active keypair
	keyPair, err := auth.LoadKeyPairFromPEM(path, kid)
	if err != nil {
		log.Fatalf("failed to load keypair: %v", err)
	}

	jwtTTL := 15 * time.Minute        // 15 min
	refreshTTL := 28 * 24 * time.Hour // 28 days

	// Build verify set - acceptable public keys
	verify := map[string]*rsa.PublicKey{
		keyPair.KID: keyPair.Public,
	}

	tokenService := auth.NewTokenService(
		keyPair,
		verify,
		jwtTTL,
		refreshTTL,
		*repositories.NewRefreshTokenRepository(queries),
	)

	return auth.ByCredentials(
		queries,
		tokenService,
	)
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
