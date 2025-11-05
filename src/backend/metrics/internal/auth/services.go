package auth

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"metrics/internal/db/sqlc"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Register(ctx context.Context, email, password string) (*authResult, error)
	Authenticate(ctx context.Context, email, password string) (*authResult, error)
}

// CredentialService handles user authentication and registration using basic email-password credentials.
// It hashes passwords using bcrypt, interacts with the database to persist user records, and generates JWT tokens
// for session management. It relies on sqlc.Queries for database operations and supports configurable JWT timeouts.
type CredentialService struct {
	queries      *sqlc.Queries
	tokenService *tokenService
}

func (s *CredentialService) Register(ctx context.Context, email, password string) (*SessionResult, error) {
	if email == "" || password == "" {
		return nil, ErrEmptyCredentials
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrHashingFailed, err)
	}

	user, err := s.queries.CreateUser(ctx, sqlc.CreateUserParams{
		Email:        email,
		PasswordHash: string(hash),
	})

	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrDatabaseError, err)
	}

	tokenPair, err := s.tokenService.generateTokenPair(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	return &SessionResult{
		UserID:    user.ID.String(),
		Email:     user.Email,
		TokenPair: tokenPair,
	}, nil
}

func (s *CredentialService) Authenticate(ctx context.Context, email, password string) (*SessionResult, error) {
	user, err := s.queries.GetUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// User doesn't exist
			return nil, ErrInvalidCredentials
		}
		// This is a database/infrastructure error, not an authentication error
		return nil, fmt.Errorf("%s: %v", ErrDatabaseError, err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return nil, ErrInvalidCredentials
	}

	tokenPair, err := s.tokenService.generateTokenPair(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	return &SessionResult{
		UserID:    user.ID.String(),
		Email:     user.Email,
		TokenPair: tokenPair,
	}, nil
}
