package auth

import (
	"context"
	"metrics/internal/models"
)

type Provider interface {
	Register(ctx context.Context, email, password string) (*Result, error)
	Authenticate(ctx context.Context, email, password string) (*Result, error)
}

type Result struct {
	UserID       string
	JWTToken     models.JWTToken
	RefreshToken models.RefreshToken
	Email        string
}

type MiddlewareConfig struct {
	AuthProvider Provider
}

func NewMiddlewareConfig(authProvider Provider) MiddlewareConfig {
	return MiddlewareConfig{AuthProvider: authProvider}
}
