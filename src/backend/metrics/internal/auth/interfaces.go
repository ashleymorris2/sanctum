package auth

import (
	"context"
)

type Provider interface {
	Register(ctx context.Context, email, password string) (*CredentialAuthResult, error)
	Authenticate(ctx context.Context, email, password string) (*CredentialAuthResult, error)
}
