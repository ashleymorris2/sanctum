package auth

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/google/uuid"
	"metrics/internal/model"
	"time"
)

func generateRefreshToken(time.Duration) (model.RefreshToken, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return model.NewRefreshToken(uuid.New(), base64.URLEncoding.EncodeToString(b)), nil
}
