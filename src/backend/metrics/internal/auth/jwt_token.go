package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"metrics/internal/model"
	"time"
)

func generateJWT(userID uuid.UUID, tokenTTL time.Duration, jwtSecret []byte) (model.JWTToken, error) {
	now := time.Now()
	claims := jwt.MapClaims{
		"sub": userID.String(),          //Subject - who is the token for
		"exp": now.Add(tokenTTL).Unix(), // Expiration time (unix timestamp)
		"iat": now.Unix(),               // Issued at: time when the token was generated (unix timestamp)
		"nbf": now.Unix(),               //Not before: defines the time before which the JWT cannot be accepted for processing.
		"jti": now.String(),             //JWT ID: an identifier for the JWT, which can be used to prevent the JWT from being replayed.
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return model.NewJWTToken(jwtToken), nil
}
