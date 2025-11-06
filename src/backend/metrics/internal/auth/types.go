package auth

import (
	"metrics/internal/model"
	"time"

	"github.com/google/uuid"
)

// EmailPasswordCredentials represents email/password authentication
type EmailPasswordCredentials struct {
	Email    string
	Password string
}

// SessionResult represents a complete authenticated session with tokens
type SessionResult struct {
	UserID    string
	Email     string
	TokenPair *TokenPair
}

// TokenPair represents a JWT access token and refresh token pair
type TokenPair struct {
	AccessToken     model.JWTToken
	AccessTokenTTL  time.Duration
	RefreshToken    model.RefreshToken
	RefreshTokenTTL time.Duration
}

// authResult represents the result of authentication (identity only)
type authResult struct {
	userID uuid.UUID
	email  string
}

// refreshTokenInfo contains validated refresh token information
type refreshTokenInfo struct {
	userID    uuid.UUID
	expiresAt time.Time
}
