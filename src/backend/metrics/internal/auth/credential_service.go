package auth

import (
	"context"
	"crypto/subtle"
	"database/sql"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"metrics/internal/db/repositories"
	"metrics/internal/db/sqlc"

	"metrics/internal/model"
	"time"
)

var (
	ErrAuthFailure            = errors.New("invalid credentials")
	ErrDatabaseFailure        = errors.New("database error during authentication")
	ErrInvalidJwtToken        = errors.New("invalid token")
	ErrJwtTokenGeneration     = errors.New("failed to generate authentication token")
	ErrRefreshTokenGeneration = errors.New("failed to generate refresh token")
	ErrInvalidRefreshToken    = errors.New("invalid refresh token")
	ErrExpiredRefreshToken    = errors.New("refresh token has expired")
)

// CredentialAuthResult represents the result of an authentication operation using credentials.
// It includes the authenticated user's ID, JWT access token, a refresh token for session renewal, and the user's email.
type CredentialAuthResult struct {
	UserID       string
	JWTToken     model.JWTToken
	RefreshToken model.RefreshToken
	Email        string
}

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

func (cs *CredentialService) Register(ctx context.Context, email, password string) (*CredentialAuthResult, error) {
	if email == "" || password == "" {
		return nil, errors.New("email and/or password cannot be empty")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	user, err := cs.queries.CreateUser(ctx, sqlc.CreateUserParams{
		Email:        email,
		PasswordHash: string(hash),
	})

	if err != nil {
		return nil, err
	}

	jwtToken, err := generateJWT(user.ID, cs.jwtTokenTTL, cs.jwtSecret)
	if err != nil {
		return nil, err
	}

	refreshToken, err := generateRefreshToken()
	if err != nil {
		//todo log
		return nil, ErrRefreshTokenGeneration
	}

	return &CredentialAuthResult{
		UserID:       user.ID.String(),
		JWTToken:     jwtToken,
		RefreshToken: refreshToken,
		Email:        user.Email,
	}, nil
}

func (cs *CredentialService) Authenticate(ctx context.Context, email, password string) (*CredentialAuthResult, error) {
	user, err := cs.queries.GetUserByEmail(ctx, email)
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

	jwtToken, err := generateJWT(user.ID, cs.jwtTokenTTL, cs.jwtSecret)
	if err != nil {
		//todo log
		return nil, ErrJwtTokenGeneration
	}

	refreshToken, err := generateRefreshToken()
	if err != nil {
		//todo log
		return nil, ErrRefreshTokenGeneration
	}

	err = cs.refreshTokenRepo.InsertRefreshToken(ctx, refreshToken, user.ID, cs.jwtTokenTTL)
	if err != nil {
		return nil, fmt.Errorf("%s: %v", ErrDatabaseFailure, err)
	}

	return &CredentialAuthResult{
		UserID:       user.ID.String(),
		JWTToken:     jwtToken,
		RefreshToken: refreshToken,
		Email:        user.Email,
	}, nil
}

func (cs *CredentialService) ValidateJwtToken(jwtToken model.JWTToken) (jwt.MapClaims, error) {

	// Parse the token
	token, err := jwt.Parse(jwtToken.String(), func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidJwtToken
		}
		return cs.jwtSecret, nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, err
		}
		return nil, ErrInvalidJwtToken
	}

	// Extract and validate claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, ErrInvalidJwtToken
}

func (cs *CredentialService) RefreshJwtToken(refreshToken model.RefreshToken) (model.JWTToken, error) {
	token, err := cs.ValidateRefreshToken(refreshToken)
	if err != nil {
		return "", err
	}

	// Get the user
	user, err := cs.queries.GetUserById(context.Background(), token.UserID)
	if err != nil {
		return "", err
	}

	// Generate a new access token
	jwtToken, err := generateJWT(user.ID, cs.jwtTokenTTL, cs.jwtSecret)
	if err != nil {
		return "", err
	}

	return jwtToken, nil
}

func (cs *CredentialService) ValidateRefreshToken(refreshToken model.RefreshToken) (*sqlc.RefreshToken, error) {
	// Retrieve the refresh token
	token, err := cs.refreshTokenRepo.GetRefreshToken(context.Background(), refreshToken)
	if err != nil {
		return nil, ErrInvalidRefreshToken
	}

	if subtle.ConstantTimeCompare([]byte(token.Token), []byte(refreshToken.Hashed())) != 1 {
		return nil, errors.New("invalid token")
	}

	// Check if the token is revoked
	if token.Revoked {
		return nil, ErrInvalidRefreshToken
	}

	// Check if the token has expired
	if time.Now().After(token.ExpiresAt) {
		return nil, ErrExpiredRefreshToken
	}

	return token, nil
}
