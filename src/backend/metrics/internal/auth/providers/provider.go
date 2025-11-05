package providers

import "context"

type AuthProvider interface {
	Authenticate(ctx context.Context, credentials Credentials) (*AuthResult, error)
	Register(ctx context.Context, credentials Credentials) (*AuthResult, error)
}

type Credentials struct {
	Email    string
	Password string
}

type AuthResult struct {
	UserID string
	Email  string
}
