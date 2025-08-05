package auth

import (
	"context"
	"metrics/internal/dto"
)

type Provider interface {
	Register(ctx context.Context, email, password string) (*dto.CredentialAuthResult, error)
	Authenticate(ctx context.Context, email, password string) (*dto.CredentialAuthResult, error)
}
