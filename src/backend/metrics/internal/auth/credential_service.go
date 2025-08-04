package auth

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"metrics/internal/db/repositories"
	"metrics/internal/db/sqlc"
	"time"
)

var (
	ErrAuthFailure            = errors.New("invalid credentials")
	ErrDatabaseFailure        = errors.New("database error during authentication")
	ErrJwtTokenGeneration     = errors.New("failed to generate authentication token")
	ErrRefreshTokenGeneration = errors.New("failed to generate refresh token")
)

// CredentialService handles user authentication and registration using basic email-password credentials.
// It hashes passwords using bcrypt, interacts with the database to persist user records, and generates JWT tokens
// for session management. It relies on sqlc.Queries for database operations and supports configurable JWT timeouts.
type CredentialService struct {
	queries          *sqlc.Queries
	refreshTokenRepo repositories.RefreshTokenRepository
	jwtSecret        []byte
	jwtTokenTTL      time.Duration
}

// Option modifies the configuration of a CredentialService by applying custom settings through functional options.
type Option func(*CredentialService)

// WithTokenTimeout sets the token expiration timeout to the specified duration in a CredentialService instance.
func WithTokenTimeout(d time.Duration) Option {
	return func(p *CredentialService) {
		p.jwtTokenTTL = d
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
func ByCredentials(queries *sqlc.Queries, refreshTokenRepo repositories.RefreshTokenRepository, jwtSecret []byte, opts ...Option) *CredentialService {
	p := &CredentialService{
		queries:          queries,
		refreshTokenRepo: refreshTokenRepo,
		jwtSecret:        jwtSecret,
		jwtTokenTTL:      15 * time.Minute,
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

	jwtToken, err := generateJWT(user.ID, p.jwtTokenTTL, p.jwtSecret)
	if err != nil {
		return nil, err
	}

	//refreshToken, err := GenerateRefreshToken()
	//if err != nil {
	//	return nil, err
	//}

	return &Result{UserID: user.ID.String(), Email: user.Email, JWTToken: jwtToken}, nil
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

	jwtToken, err := generateJWT(user.ID, p.jwtTokenTTL, p.jwtSecret)
	if err != nil {
		//todo log
		return nil, ErrJwtTokenGeneration
	}

	refreshToken, err := GenerateRefreshToken()
	if err != nil {
		//todo log
		return nil, ErrRefreshTokenGeneration
	}

	err = p.refreshTokenRepo.InsertRefreshToken(ctx, refreshToken, user.ID, p.jwtTokenTTL)
	if err != nil {
		return nil, fmt.Errorf("%s: %v", ErrDatabaseFailure, err)
	}

	return &Result{
		UserID:       user.ID.String(),
		JWTToken:     jwtToken,
		RefreshToken: refreshToken,
		Email:        user.Email,
	}, nil
}
