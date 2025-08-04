package auth

import (
	"crypto/rand"
	"encoding/base64"
	"metrics/internal/models"
)

func GenerateRefreshToken() (models.RefreshToken, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return models.NewRefreshToken(base64.URLEncoding.EncodeToString(b)), nil
}
