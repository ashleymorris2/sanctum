package auth

import (
	"context"
	"metrics/internal/model"

	"github.com/google/uuid"
)

// authProvider is the base interface that all providers implement
type authProvider interface {
	// name returns the provider type (for debugging/logging)
	name() string
}

type credentialAuthProvider interface {
	authProvider
	authenticateWithCredentials(ctx context.Context, credentials EmailPasswordCredentials) (*authResult, error)
	registerWithCredentials(ctx context.Context, credentials EmailPasswordCredentials) (*authResult, error)
}

// CredentialService handles email/password authentication
type CredentialService interface {
	Login(ctx context.Context, credentials EmailPasswordCredentials) (*SessionResult, error)
	Register(ctx context.Context, credentials EmailPasswordCredentials) (*SessionResult, error)
	RefreshSession(ctx context.Context, refreshToken model.RefreshToken) (*TokenPair, error)
	ValidateToken(token model.JWTToken) (uuid.UUID, error)
	Logout(ctx context.Context, refreshToken model.RefreshToken) error
}
