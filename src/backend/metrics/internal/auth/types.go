package auth

import (
	"metrics/internal/model"
	"time"

	"github.com/google/uuid"
)

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
	UserID string
	Email  string
}

// refreshTokenInfo contains validated refresh token information
type refreshTokenInfo struct {
	UserID    uuid.UUID
	ExpiresAt time.Time
}
