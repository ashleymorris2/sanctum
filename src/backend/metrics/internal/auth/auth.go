package auth

import "context"

type Provider interface {
	Register(ctx context.Context, email, password string) (*Result, error)
	Authenticate(ctx context.Context, email, password string) (*Result, error)
}

type Result struct {
	UserID string
	Token  string
	Email  string
}
