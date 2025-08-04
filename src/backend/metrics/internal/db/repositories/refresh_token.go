package repositories

import (
	"context"
	"github.com/google/uuid"

	"metrics/internal/models"

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

func (r *RefreshTokenRepository) InsertRefreshToken(ctx context.Context, refreshToken models.RefreshToken, userID uuid.UUID, ttl time.Duration) error {
	tokenID := uuid.New()
	expiresAt := time.Now().Add(ttl)
	params := &sqlc.InsertRefreshTokenParams{
		ID:        tokenID,
		UserID:    userID,
		Token:     refreshToken.String(),
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
