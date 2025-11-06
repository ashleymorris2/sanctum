package auth

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"metrics/internal/db/sqlc"

	"golang.org/x/crypto/bcrypt"
)

type credentialProvider struct {
	queries *sqlc.Queries
}

func newCredentialProvider(queries *sqlc.Queries) *credentialProvider {
	return &credentialProvider{queries: queries}
}

func (p *credentialProvider) name() string {
	return "credentials"
}

func (p *credentialProvider) authenticateWithCredentials(ctx context.Context, creds EmailPasswordCredentials) (*authResult, error) {
	if creds.Email == "" || creds.Password == "" {
		return nil, ErrEmptyCredentials
	}

	user, err := p.queries.GetUserByEmail(ctx, creds.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrInvalidCredentials
		}
		return nil, fmt.Errorf("%w: %v", ErrDatabaseError, err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(creds.Password)); err != nil {
		return nil, ErrInvalidCredentials
	}

	return &authResult{
		userID: user.ID,
		email:  user.Email,
	}, nil
}

func (p *credentialProvider) registerWithCredentials(ctx context.Context, creds EmailPasswordCredentials) (*authResult, error) {
	if creds.Email == "" || creds.Password == "" {
		return nil, ErrEmptyCredentials
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(creds.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrHashingFailed, err)
	}

	user, err := p.queries.CreateUser(ctx, sqlc.CreateUserParams{
		Email:        creds.Email,
		PasswordHash: string(hash),
	})
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrDatabaseError, err)
	}

	return &authResult{
		userID: user.ID,
		email:  user.Email,
	}, nil
}
