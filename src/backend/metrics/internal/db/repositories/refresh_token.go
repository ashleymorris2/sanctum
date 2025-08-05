package repositories

import (
	"context"
	"errors"
	"github.com/google/uuid"

	"metrics/internal/model"

	"metrics/internal/db/sqlc"
	"time"
)

// RefreshTokenRepository handles database operations for refresh tokens
type RefreshTokenRepository struct {
	db *sqlc.Queries
}

func NewRefreshTokenRepository(db *sqlc.Queries) *RefreshTokenRepository {
	return &RefreshTokenRepository{db: db}
}

func (r *RefreshTokenRepository) InsertRefreshToken(ctx context.Context, refreshToken model.RefreshToken, userID uuid.UUID, ttl time.Duration) error {
	expiresAt := time.Now().Add(ttl)
	params := &sqlc.InsertRefreshTokenParams{
		ID:        refreshToken.ID(),
		UserID:    userID,
		Token:     refreshToken.Hashed(),
		ExpiresAt: expiresAt,
		CreatedAt: time.Now(),
		Revoked:   false,
	}

	_, err := r.db.InsertRefreshToken(ctx, *params)
	if err != nil {
		return err
	}
	return nil
}

func (r *RefreshTokenRepository) GetRefreshToken(ctx context.Context, refreshToken model.RefreshToken) (*sqlc.RefreshToken, error) {
	token, err := r.db.GetRefreshTokenById(ctx, refreshToken.ID())
	if err != nil {
		return nil, errors.New("token not found")
	}

	return &token, nil
}

func (r *RefreshTokenRepository) InvalidateRefreshToken(token model.RefreshToken) error {
	err := r.db.InvalidateRefreshTokenById(context.Background(), token.ID())
	if err != nil {
		return errors.New("error invalidating token")
	}

	return nil
}
