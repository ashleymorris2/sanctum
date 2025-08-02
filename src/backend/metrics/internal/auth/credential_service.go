package auth

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"metrics/internal/db/sqlc"
	"time"
)

var (
	ErrAuthFailure     = errors.New("invalid credentials")
	ErrDatabaseFailure = errors.New("database error during authentication")
	ErrTokenGeneration = errors.New("failed to generate authentication token")
)

// CredentialService handles user authentication and registration using basic email-password credentials.
// It hashes passwords using bcrypt, interacts with the database to persist user records, and generates JWT tokens
// for session management. It relies on sqlc.Queries for database operations and supports configurable JWT timeouts.
type CredentialService struct {
	queries      *sqlc.Queries
	jwtSecret    []byte
	tokenTimeout time.Duration
}

// Option modifies the configuration of a CredentialService by applying custom settings through functional options.
type Option func(*CredentialService)

// WithTokenTimeout sets the token expiration timeout to the specified duration in a CredentialService instance.
func WithTokenTimeout(d time.Duration) Option {
	return func(p *CredentialService) {
		p.tokenTimeout = d
	}
}

// ByCredentials creates a new instance of CredentialService with the specified database
// queries and JWT secret.
// It configures a default token timeout of 15 minutes, which can be
// modified using functional options.
//
// Example:
//
//	// Create a provider with custom token timeout
//	provider := auth.ByCredentials(
//	    queries,
//	    []byte("your-jwt-secret"),
//	    auth.WithTokenTimeout(24 * time.Hour),
//	)
func ByCredentials(queries *sqlc.Queries, jwtSecret []byte, opts ...Option) *CredentialService {
	p := &CredentialService{
		queries:      queries,
		jwtSecret:    jwtSecret,
		tokenTimeout: 15 * time.Minute,
	}
	for _, opt := range opts {
		opt(p)
	}
	return p
}

func (p *CredentialService) Register(ctx context.Context, email, password string) (*Result, error) {
	if email == "" || password == "" {
		return nil, errors.New("email and/or password cannot be empty")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	user, err := p.queries.CreateUser(ctx, sqlc.CreateUserParams{
		Email:        email,
		PasswordHash: string(hash),
	})

	if err != nil {
		return nil, err
	}

	token, err := p.generateJWT(user.ID)
	if err != nil {
		return nil, err
	}

	return &Result{UserID: user.ID.String(), Email: user.Email, Token: token}, nil
}

func (p *CredentialService) Authenticate(ctx context.Context, email, password string) (*Result, error) {
	user, err := p.queries.GetUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// User doesn't exist
			return nil, ErrAuthFailure
		}
		// This is a database/infrastructure error, not an authentication error
		return nil, fmt.Errorf("%s: %v", ErrDatabaseFailure, err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return nil, ErrAuthFailure
	}

	token, err := p.generateJWT(user.ID)
	if err != nil {
		return nil, ErrTokenGeneration
	}

	return &Result{UserID: user.ID.String(), Email: user.Email, Token: token}, nil
}

func (p *CredentialService) generateJWT(userID uuid.UUID) (string, error) {
	now := time.Now()
	claims := jwt.MapClaims{
		"sub": userID.String(),                //Subject - who is the token for
		"exp": now.Add(p.tokenTimeout).Unix(), // Expiration time (unix timestamp)
		"iat": now.Unix(),                     // Issued at: time when the token was generated (unix timestamp)
		"nbf": now.Unix(),                     //Not before: defines the time before which the JWT cannot be accepted for processing.
		"jti": now.String(),                   //JWT ID: an identifier for the JWT, which can be used to prevent the JWT from being replayed.
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(p.jwtSecret)
}
