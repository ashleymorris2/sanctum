package auth

import (
	"metrics/internal/model"
	"time"

	"github.com/google/uuid"
)

// CredentialAuthResult represents the result of an authentication operation using credentials.
// It includes the authenticated user's ID, JWT access token, a refresh token for session renewal, and the user's email.
type CredentialAuthResult struct {
	UserID    string
	Email     string
	TokenPair *TokenPair
}

// RefreshTokenInfo contains validated refresh token information
type RefreshTokenInfo struct {
	UserID    uuid.UUID
	ExpiresAt time.Time
}

// TokenPair represents a JWT access token and refresh token pair
type TokenPair struct {
	AccessToken     model.JWTToken
	AccessTokenTTL  time.Duration
	RefreshToken    model.RefreshToken
	RefreshTokenTTL time.Duration
}
